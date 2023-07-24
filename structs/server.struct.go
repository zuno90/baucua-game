package structs

import (
	"fmt"
	"reflect"
)


type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Rooms      map[string]*Room
	Broadcast  chan Payload
	Action chan Payload
}

// singleton pattern - create server/lobby room
func ServerInstance() *Server {
	return &Server{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Rooms:      make(map[string]*Room),
		Broadcast:  make(chan Payload),
		Action: make(chan Payload),
	}
}

func (server *Server) ListenChannel() {
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
			if message.Type == "CHAT" {
				// broadcast to all
				mv := reflect.ValueOf(message)
				// fmt.Println(mv.FieldByName(message.Msg).IsValid())
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

		case ma := <-server.Action:
			if ma.Type ==  "CREATEROOM" {
				host := server.Clients[ma.From]
				if err := host.CreateRoom(); err != nil {
					ms := Payload{From: ma.From, Msg: "Bad request! Can not create a new room!"}
					host.Send(ms)
				}
				
			}
		}
	}
}
