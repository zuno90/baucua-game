package structs

import (
	"fmt"
	"reflect"
)

type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Match      chan *Client
	Broadcast  chan Payload
}

// singleton pattern
func ServerInstance() *Server {
	return &Server{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Match:      make(chan *Client),
		Broadcast:  make(chan Payload),
	}
}

func (server *Server) Start() {
	for {
		select {
		case nc := <-server.Register:
			server.Clients[nc.ID] = nc
			fmt.Println("Size of Connection Server: ", len(server.Clients))
			mess := Payload{
				From: nc.ID,
				Msg:  fmt.Sprintf("User %s Connected!...", nc.ID),
			}
			fmt.Println(mess)
			for _, client := range server.Clients {
				if nc.ID != client.ID {
					client.Send(mess)
				}
			}

		case ec := <-server.Unregister:
			delete(server.Clients, ec.ID)
			fmt.Println("Size of Connection Server: ", len(server.Clients))
			mess := Payload{
				From: ec.ID,
				Msg:  fmt.Sprintf("User %s Disconnected!...", ec.ID),
			}
			for _, client := range server.Clients {
				if ec.ID != client.ID {
					client.Send(mess)
				}
			}

		case message := <-server.Broadcast:
if message.Type == "chat" {
				// broadcast to all
				mv := reflect.ValueOf(message)
				if message.To == "all" || !mv.FieldByName(message.To).IsValid() {
					for _, client := range server.Clients {
						if client.ID != message.From {
							client.Send(message)
						}
					}
				}
				// Pm to user
				if rc, ok := server.Clients[message.To]; ok {
					if rc.ID == message.To {
						rc.Send(message)
					}
				}
			}
		}
	}
}
