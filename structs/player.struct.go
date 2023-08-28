package structs

import "log"

type PlayerAction interface {
	joinRoom() error
	leaveRoom() error
}

type Player struct {
	Id     int32   `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func NewPlayer(id int32, username string, amount float64) *Player {
	return &Player{
		Id:     id,
		Name:   username,
		Amount: amount,
	}
}

func (p *Player) joinRoom() {
	log.Printf("player %s has joined!", p.Name)
}
