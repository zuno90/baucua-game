package handlers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	pb "github.com/zuno90/go-ws/pb"
	st "github.com/zuno90/go-ws/structs"
)

// var (
// 	msgType int
// 	msg     []byte
// 	err     error
// )

func HandleConn(c *websocket.Conn, s *st.Server, user *pb.User) {
	uniqueID := uuid.New()
	np := st.NewPlayer(user.GetId(), user.GetUsername(), user.GetAmount())

	client := &st.Client{
		ID:     uniqueID.String(),
		Conn:   c,
		Server: s,
		Room:   make([]string, 1),
		Player: make(map[string]*st.Player),
	}
	client.Player[client.ID] = np
	client.ConnectToServer()
}
