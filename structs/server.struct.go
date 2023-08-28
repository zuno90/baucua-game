package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Rooms      map[string]*Room
	Broadcast  chan Payload
	Action     chan Payload
}

// singleton pattern - create server/lobby room
func ServerInstance() *Server {
	return &Server{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Rooms:      make(map[string]*Room),
		Broadcast:  make(chan Payload),
		Action:     make(chan Payload),
	}
}

func (server *Server) ListenChannel() {
	for {
		select {
		case nc := <-server.Register:
			// send info to Client
			clientInfo, err := json.Marshal(nc.Player)
			if err != nil {
				fmt.Println("Payload is wrong format :::::", err)
				return
			}
			info := ResMessage("LOGIN", string(clientInfo))
			nc.Send(info)
			server.Clients[nc.ID] = nc
			fmt.Println("Size of Connection Server (connect): ", len(server.Clients))
			for _, client := range server.Clients {
				if nc.ID != client.ID {
					mess := ResMessage("WELCOME", fmt.Sprintf("User %s Connected!...", nc.ID))
					client.Send(mess)
				}
			}

		case ec := <-server.Unregister:
			delete(server.Clients, ec.ID)
			fmt.Println("Size of Connection Server (disconnect): ", len(server.Clients))
			for _, client := range server.Clients {
				if ec.ID != client.ID {
					mess := ResMessage("BYEBYE", fmt.Sprintf("User %s Disconnected!...", ec.ID))
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
							log.Println("message broadcast", message)
							// client.Send(message)
						}
					}
				}
				// Pm to user
				if rc, ok := server.Clients[message.To]; ok {
					if rc.ID == message.To {
						log.Println("message broadcast", message)
						// rc.Send(message)
					}
				}
			}

			// case ma := <-server.Action:
			// 	client := server.Clients[ma.From]
			// 	switch ma.Type {
			// 	case "CREATEROOM":
			// 		fmt.Println("client in room", client.Room)
			// 		fmt.Println("length of client in room", len(client.Room))
			// 		if len(client.Room) > 0 {
			// 			em := ResMessage(ma.From, "You was currently in Room!")
			// 			client.Send(em)
			// 		}
			// 		if err := client.CreateRoom(); err != nil {
			// 			ms := ResMessage(ma.From, "Bad request! Can not create a new room!")
			// 			client.Send(ms)
			// 		}
			// 	case "JOINROOM":
			// 		roomId := ma.Msg
			// 		r, ok := server.Rooms[roomId]
			// 		if !ok {
			// 			ms := ResMessage(ma.From, "Room is not existing!")
			// 			client.Send(ms)
			// 		}
			// 		if err := client.JoinRoom(r); err != nil {
			// 			log.Fatal(err)
			// 		}
			// 	}
		}
	}
}
