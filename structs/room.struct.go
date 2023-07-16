package structs

import (
	"fmt"
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
	Join chan *Player
}

func (room *Room) ListenChannel() {
	for {
		select {
		case jp := <- room.Join:
			fmt.Println("new player had joined room", jp)

		}
	}
}