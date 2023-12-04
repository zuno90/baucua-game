package game

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/zuno90/go-ws/utils"
)

type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Broadcast  chan ResData
	Action     chan ResData
}

// singleton pattern - create server/lobby room
func ServerInstance() *Server {
	return &Server{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan ResData),
		Action:     make(chan ResData),
	}
}

func (server *Server) StartGame() {
	log.Println("start game...")
	fn := func() ResData {
		stage, gameno, sec := HandleGameInfo()
		if stage == Stages(RESULTTIME) && sec == 5 {
			server.HandleResult()
		}
		// countdown timer clock
		ng := server.NewGame(stage, gameno, sec)
		gameinfo, err := json.Marshal(ng)
		if err != nil {
			log.Println("Can not marshal :::::", err)
		}
		return ResMessage(Types(GAMEINFO), string(gameinfo))
	}
	go func() {
		for {
			server.Broadcast <- fn()
			time.Sleep(time.Second)
		}
	}()
}

func (server *Server) HandleResult() {
	result := utils.RandomResult()
	rm := ResMessage(Types(RESULT), strings.Join(result, ","))
	go func() {
		server.Broadcast <- rm
	}()
}

func (server *Server) ListenEvents() {
	for {
		select {
		case nc := <-server.Register:
			// send info to Client
			clientInfo, err := json.Marshal(nc.Player)
			if err != nil {
				log.Println("Can not marshal :::::", err)
				server.Unregister <- nc
			}

			// client info (private)
			info := ResMessage(Types(LOGIN), string(clientInfo))
			fmt.Println("vao day", info)
			nc.Send(info)

			// inform for others users (public)
			server.Clients[nc.ID] = nc
			log.Println("Size of Connection Server (connect): ", len(server.Clients))
			for _, client := range server.Clients {
				if nc.ID != client.ID {
					mess := ResMessage(Types(WELCOME), fmt.Sprintf("%s Connected!...", nc.Player[nc.ID].Name))
					client.Send(mess)
				}
			}

		case ec := <-server.Unregister:
			delete(server.Clients, ec.ID)
			log.Println("Size of Connection Server (disconnect): ", len(server.Clients))
			for _, client := range server.Clients {
				if ec.ID != client.ID {
					mess := ResMessage(Types(BYEBYE), fmt.Sprintf("%s Disconnected!...", ec.Player[ec.ID].Name))
					client.Send(mess)
				}
			}

		case message := <-server.Broadcast:
			switch message.Type {
			case Types(CHAT):
				fmt.Println(message)
				mv := reflect.ValueOf(message)
				// log.Println(mv.FieldByName(message.Msg).IsValid())
				if message.To == "all" || !mv.FieldByName(message.To).IsValid() {
					for _, client := range server.Clients {
						if client.ID != message.From {
							log.Println("message broadcast", message)
							client.Send(message)
						}
					}
				}
				// Pm to user
				if rc, ok := server.Clients[message.To]; ok {
					if rc.ID == message.To {
						log.Println("message broadcast", message)
						rc.Send(message)
					}
				}

			case Types(GAMEINFO):
				for _, client := range server.Clients {
					client.Send(message)
				}

			case Types(RESULT):
				log.Println("game random result", message)
				for _, client := range server.Clients {
					client.Send(message)
				}
			}
		case ma := <-server.Action:
			fmt.Printf("%+v\n", ma)
		}
	}

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
	// }
}
