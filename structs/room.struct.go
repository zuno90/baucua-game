package structs

import (
	"github.com/google/uuid"
)

type Levels string

const (
	LOW  Levels = "LOW"
	MID  Levels = "MID"
	HIGH Levels = "HIGH"
)

type Room struct {
	ID      string
	Level   Levels
	Players map[string]*Player
	Join    chan *Player
}

func (c *Client) RoomInstance(l Levels) *Room {
	uniqueId := uuid.New()
	return &Room{
		ID:      uniqueId.String(),
		Level:   l,
		Players: make(map[string]*Player),
		Join:    make(chan *Player),
	}
}

func (room *Room) ListenChannel() {
	for {
		select {
		case jp := <-room.Join:
			// fmt.Println("new player had joined room", jp)
			// fmt.Println("room state", room)
			// fmt.Println("player in room", *jp)
			jp.joinRoom()
		
		// case <-room.Leave:
		
		}

	}
}
