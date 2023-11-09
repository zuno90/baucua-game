package game

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"golang.org/x/exp/slices"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Server *Server
	Room   []string
	Player map[string]*Player
}

type Joiner struct {
	CId string
	PId int32
}

var players []*Joiner

func (c *Client) ConnectToServer() error {
	defer func() {
		// remove id from players slice
		if idx := slices.IndexFunc(players, func(j *Joiner) bool { return j.CId == c.ID }); idx > -1 {
			players = slices.Delete(players, idx, idx+1)
		}
		// disconnect & close connection
		c.Server.Unregister <- c
		c.Conn.Close()
	}()
	// validate player existing player id
	for _, p := range players {
		if c.Player[c.ID].Id == p.PId {
			c.SendError(ResErrorMessage("ERROR", 403, "Player is existing!"))
			return fmt.Errorf("Player is existing!")
		}
	}
	j := &Joiner{CId: c.ID, PId: c.Player[c.ID].Id}
	players = append(players, j)
	c.Server.Register <- c

	// listen all events
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Bad connection :::::", err)
			return err
		}
		nm := ResData{From: c.ID}
		if err := json.Unmarshal([]byte(msg), &nm); err != nil {
			log.Println("Payload is wrong format :::::", err)
			m := ResErrorMessage(nm.Type, 400, "Payload is wrong format") /* Payload{From: c.ID, Msg: "Payload is wrong format!"} */
			c.SendError(m)
		}
		log.Println("wkdfjwdef", nm)
		switch nm.Type {
		case Types(CHAT):
			c.Server.Broadcast <- nm
		default:
			log.Println("xuong dayyyyy")
			m := ResErrorMessage(nm.Type, 400, "Payload is wrong format") /* Payload{From: c.ID, Msg: "Payload is wrong format!"} */
			c.SendError(m)
		}
	}
}

func (c *Client) Send(d ResData) error {
	resD, err := json.Marshal(d)
	if err != nil {
		log.Println("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resD); err != nil {
		return err
	}
	return nil
}

func (c *Client) SendError(e ResError) error {
	resE, err := json.Marshal(e)
	if err != nil {
		log.Println("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resE); err != nil {
		return err
	}
	return nil
}
