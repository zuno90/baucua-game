package structs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() error {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Bad connection :::::",err)
			return err
		}
		nm := Payload{From: c.ID}
		nm.Type = "chat"
		if err := json.Unmarshal([]byte(msg), &nm); err != nil {
			fmt.Println("Payload is wrong format :::::",err)
			m := Payload{From: c.ID, Msg: "Payload is wrong format!"}
			c.Send(m)
		}
		c.Pool.Broadcast <- nm
	}
}

func (c *Client) Send(payload Payload) error {
	if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(payload.Msg)); err != nil {
		return err
	}
	return nil
}
