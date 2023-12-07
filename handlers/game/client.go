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

// var players []*Joiner

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
			log.Println("Bad connection :::::", err)
			return err
		}
		nm := ResData{From: c.ID}
		if err := json.Unmarshal([]byte(msg), &nm); err != nil {
			log.Println("Payload is wrong format :::::", err)
			m := ResErrorMessage(nm.Type, 400, "Payload is wrong format") /* Payload{From: c.ID, Msg: "Payload is wrong format!"} */
			c.SendError(m)
		}
		// client action by type
		switch nm.Type {
		case Types(CHAT):
			c.Server.Broadcast <- nm
		case Types(BET):
			c.Server.Action <- nm
		default:
			m := ResErrorMessage(nm.Type, 400, "Payload is wrong format") /* Payload{From: c.ID, Msg: "Payload is wrong format!"} */
			c.SendError(m)
		}
	}
}

func (c *Client) Send(d ResData) error {
	resD, err := json.Marshal(d)
	if err != nil {
		log.Println("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resD); err != nil {
		return err
	}
	return nil
}

func (c *Client) SendError(e ResError) error {
	resE, err := json.Marshal(e)
	if err != nil {
		log.Println("Can not marshal :::", err)
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, resE); err != nil {
		return err
	}
	return nil
}

func (c *Client) addPlayer() error {
	for k, v := range c.Player {
		// validate player existing player id
		pkey := fmt.Sprintf("players:%d", k)
		p, _ := utils.Get(pkey)
		if p != nil {
			c.SendError(ResErrorMessage("ERROR", 403, "Player is existing!"))
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
	fmt.Println("remove player")
	for k := range c.Player {
		pkey := fmt.Sprintf("players:%d", k)
		if err := utils.Del(pkey); err != nil {
			fmt.Sprintln("can not remove player ", k)
		}
	}
}
