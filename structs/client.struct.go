package structs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Server *Server
	Room *Room
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
		// nm.Type = "CHAT"
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
	uniqueId := uuid.New()
	newRoom := &Room{
		ID: uniqueId.String(),
		Level: "LOW",
		Players: make(map[string]*Player),
		Join: make(chan *Player),
	}
	if err := c.JoinRoom(newRoom.ID); err != nil {
		return err
	}
	return nil
}

// join room
func (c *Client) JoinRoom(id string) error {
	// create test player
	newPlayer := c.NewPlayer("1","zuno", 100.67)
	c.Room.Join <- newPlayer
}
