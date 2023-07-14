package handlers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/zuno90/go-ws/structs"
)

// var (
// 	msgType int
// 	msg     []byte
// 	err     error
// )

func HandleConn(c *websocket.Conn, s *structs.Server) {
	uniqueID := uuid.New()
	client := &structs.Client{
		ID:     uniqueID.String(),
		Conn:   c,
		Server: s,
	}
	s.Register <- client
	client.Read()
}
