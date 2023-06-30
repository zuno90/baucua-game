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
	Pm   chan string
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		nm := Payload{From: c.ID}
		if err := json.Unmarshal([]byte(msg), &nm); err != nil {
			fmt.Println(err)
			return
		}
		c.Pool.Broadcast <- nm
	}
}

func (c *Client) Send(payload Payload) error {
	return c.Conn.WriteMessage(websocket.TextMessage, []byte(payload.Msg))
}
