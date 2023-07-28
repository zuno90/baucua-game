package structs

import "fmt"

type PlayerAction interface {
	joinRoom() error
	leaveRoom() error
}

type Player struct {
	ID   string
	Name string
	Coin float64
}

func (c *Client) NewPlayer(id string, name string, coin float64) *Player {
	return &Player{
		ID:   id,
		Name: name,
		Coin: coin,
	}
}

func (p Player) joinRoom() {
	fmt.Println("player", p.Name, "join room!")
}
