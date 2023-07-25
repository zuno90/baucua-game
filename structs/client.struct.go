package structs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Server *Server
	Room   []*Room
}

func (c *Client) ConnectToServer() error {
	defer func() {
		c.Server.Unregister <- c
		c.Conn.Close()
	}()

	// listen all events
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Bad connection :::::", err)
			return err
		}
		nm := Payload{From: c.ID}
		if err := json.Unmarshal([]byte(msg), &nm); err != nil {
			fmt.Println("Payload is wrong format :::::", err)
			m := Payload{From: c.ID, Msg: "Payload is wrong format!"}
			c.Send(m)
		}
		switch nm.Type {
		case "CHAT":
			c.Server.Broadcast <- nm
		default:
			c.Server.Action <- nm
		}
	}
}

func (c *Client) Send(payload Payload) error {
	if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(payload.Msg)); err != nil {
		return err
	}
	return nil
}

// ROOM
// create room
func (c *Client) CreateRoom() error {
	newRoom := c.RoomInstance("LOW")
	c.Server.Rooms[newRoom.ID] = newRoom
	go newRoom.ListenChannel()
	if err := c.JoinRoom(newRoom); err != nil {
		return err
	}
	fmt.Println("room id", newRoom.ID)
	return nil
}

// join room
func (c *Client) JoinRoom(r *Room) error {
	// create test player
	newPlayer := c.NewPlayer(c.ID[0:5], "zuno" + c.ID[0:5], 1000)
	c.Room = append(c.Room, r)
	r.Players[c.ID] = newPlayer
	go func() {
		r.Join <- newPlayer
	}()
	return nil
}
