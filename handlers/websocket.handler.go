package handlers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	st "github.com/zuno90/go-ws/structs"
)

// var (
// 	msgType int
// 	msg     []byte
// 	err     error
// )

func HandleConn(c *websocket.Conn, s *st.Server) {
	uniqueID := uuid.New()
	client := &st.Client{
		ID:     uniqueID.String(),
		Conn:   c,
		Server: s,
		Room:   make([]string, 1),
	}
	s.Register <- client
	client.ConnectToServer()
}
