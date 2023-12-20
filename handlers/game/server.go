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
			infoData := map[string]interface{}{"sid": nc.ID, "user": nc.Player}
			// send info to Client
			clientInfo, err := json.Marshal(infoData)
			if err != nil {
				log.Println("Can not marshal :::::", err)
				server.Unregister <- nc
			}
			// client info (private)
			info := ResMessage(Types(LOGIN), string(clientInfo))
			fmt.Println(info, "info user")
			// im := ResData{From: nc.ID, Msg: info}
			nc.Send(info)

			// inform for others users (public)
			server.Clients[nc.ID] = nc
			log.Println("Size of Connection Server (connect): ", len(server.Clients))
			for _, v := range nc.Player {
				for _, client := range server.Clients {
					mess := ResMessage(Types(WELCOME), fmt.Sprintf("%s has connected!...", v.Name))
					client.Send(mess)
				}
			}

		case ec := <-server.Unregister:
			delete(server.Clients, ec.ID)
			for _, client := range server.Clients {
				for _, v := range ec.Player {
					mess := ResMessage(Types(BYEBYE), fmt.Sprintf("%s has disconnected!...", v.Name))
					client.Send(mess)
				}
			}
			log.Println("Size of Connection Server (disconnect): ", len(server.Clients))

		case message := <-server.Broadcast:
			switch message.Type {
			case Types(CHAT):
				mv := reflect.ValueOf(message)
				if message.To == "all" || !mv.FieldByName(message.To).IsValid() {
					for _, client := range server.Clients {
						log.Println("message broadcast all users", message)
						client.Send(message)
					}
				}

				// Pm to user
				if rc, ok := server.Clients[message.To]; ok {
					if rc.ID == message.To {
						log.Println("message broadcast to 1 user", message)
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
			// handle logic bet
			fmt.Printf("action payload %+v\n", ma)
		}
	}
}
