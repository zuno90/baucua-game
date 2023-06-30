package handlers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/zuno90/go-ws/structs"
)

var (
	msgType int
	msg     []byte
	err     error
)

func HandleConn(p *structs.Pool, c *websocket.Conn) {
	uniqueID := uuid.New()
	client := &structs.Client{
		ID:   uniqueID.String(),
		Conn: c,
		Pool: p,
	}
	p.Register <- client
	client.Read()
}
