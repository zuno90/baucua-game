package structs

type Player struct {
	ID   string
	Name string
	Coin float64
	Client *Client
}

func (c *Client) NewPlayer(id string, name string, coin float64) *Player {
	return &Player{
		ID:   id,
		Name: name,
		Coin: coin,
	}
}