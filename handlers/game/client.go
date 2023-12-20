package game

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/zuno90/go-ws/utils"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Server *Server
	Room   []string
	Player map[int32]*Player
}

var FormatError = ResErrorMessage(Types(ERROR), 400, "Payload is wrong format") // Payload{From: c.ID, Msg: "Payload is wrong format!"}

func (c *Client) ConnectToServer() error {
	defer func() {
		// if idx := slices.IndexFunc(players, func(j *Joiner) bool { return j.CId == c.ID }); idx > -1 {
		// 	players = slices.Delete(players, idx, idx+1)
		// }
		// disconnect & close connection
		c.removePlayer()
		c.Server.Unregister <- c
		c.Conn.Close()
		fmt.Println("close connection websocket!")
	}()

	// validate then add to keydb
	if err := c.addPlayer(); err != nil {
		return err
	}
	c.Server.Register <- c
	// listen all events
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Bad connection :::::", err)
			return err
		}
		for _, v := range c.Player {
			nm := ResData{From: c.ID, Name: v.Name}
			if e := json.Unmarshal([]byte(msg), &nm); e != nil || len(nm.Msg) == 0 {
				log.Println("Payload is wrong format", e)
				c.SendError(FormatError)
			} else {
				// client action by type
				switch nm.Type {
				case Types(CHAT):
					c.Server.Broadcast <- nm
				case Types(BET):
					c.Server.Action <- nm
				default:
					return fmt.Errorf("no action matched or wrong format")
				}
			}
		}
	}
}

func (c *Client) Send(d ResData) {
	resD, err := json.Marshal(d)
	if err != nil {
		log.Fatal("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resD); err != nil {
		log.Fatal("Can not send message", err)
	}
}

func (c *Client) SendError(e ResError) {
	resE, err := json.Marshal(e)
	if err != nil {
		log.Fatal("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resE); err != nil {
		log.Fatal("Can not send message", err)
	}
}

func (c *Client) addPlayer() error {
	for k, v := range c.Player {
		// validate player existing player id
		pkey := fmt.Sprintf("players:%d", k)
		p, _ := utils.Get(pkey)
		if p != nil {
			c.SendError(ResErrorMessage(Types(ERROR), 403, "Player is existing!"))
			return fmt.Errorf("Player %d is logging in, please log out first or login by other accounts!", k)
		}
		// set connected user to cache keydb
		val, err := utils.MarshalBinary(v)
		if err != nil {
			log.Fatal("can not marshal", err)
			return err
		}
		if err := utils.Set(pkey, val); err != nil {
			log.Fatal("can not set to keydb", err)
			return err
		}
	}
	return nil
}

func (c *Client) removePlayer() {
	for k := range c.Player {
		pkey := fmt.Sprintf("players:%d", k)
		if err := utils.Del(pkey); err != nil {
			fmt.Sprintln("can not remove player ", k)
		}
	}
}
