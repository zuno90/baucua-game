package structs

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Match      chan *Client
	Broadcast  chan Payload
	Pm         chan Payload
}

// singleton pattern
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Match:      make(chan *Client),
		Broadcast:  make(chan Payload),
		Pm:         make(chan Payload),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case nc := <-pool.Register:
			pool.Clients[nc.ID] = nc
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			mess := Payload{
				From: nc.ID,
				To:   "all",
				Msg:  fmt.Sprintf("User %s Connected!...", nc.ID),
			}
			fmt.Println(mess)
			for _, client := range pool.Clients {
				if nc.ID != client.ID {
					if err := client.Send(mess); err != nil {
						fmt.Println(err)
						return
					}
				}
			}

		case ec := <-pool.Unregister:
			delete(pool.Clients, ec.ID)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			mess := Payload{
				From: ec.ID,
				To:   "all",
				Msg:  fmt.Sprintf("User %s Disconnected!...", ec.ID),
			}
			for _, client := range pool.Clients {
				if ec.ID != client.ID {
					if err := client.Send(mess); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
			
		case message := <-pool.Broadcast:
			if message.To == "all" {
				for _, client := range pool.Clients {
					if client.ID != message.From {
						if err := client.Send(message); err != nil {
							fmt.Println(err)
							return
						}
					}
				}
			}
			if rc, ok := pool.Clients[message.To]; ok {
				if rc.ID == message.To {
					if err := rc.Send(message); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}
