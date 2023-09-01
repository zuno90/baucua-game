package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	hdl "github.com/zuno90/go-ws/handlers"
	grpcclient "github.com/zuno90/go-ws/handlers/grpc-client"
	pb "github.com/zuno90/go-ws/pb"
	st "github.com/zuno90/go-ws/structs"
)

var (
	user    *pb.User
	Players []int
)

func SetUpWebsocket(app *fiber.App) {
	// initialize server
	server := st.ServerInstance()
	server.StartGame()
	go server.ListenEvents()
	// websocket
	h := func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			// handle auth
			jwt := c.Query("jwt")
			u, err := grpcclient.GetAuthUser(jwt)
			if err != nil || u == nil {
				return fiber.ErrForbidden
			}
			// check if client is exist
			playerChannel := make(chan int, 1)

			go func() {
				playerChannel <- int(u.GetId())
			}()
			user = u
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
	app.Use("/ws", h)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// log.Println("jwt", c.Query("jwt")) // 1.0
		hdl.HandleConn(c, server, user)
	}))

	// app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
	// log.Println(c.Locals("allowed"))  // true
	// log.Println(c.Params("id"))       // 123
	// log.Println(c.Query("v"))         // 1.0
	// log.Println(c.Cookies("session")) // ""

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	// hdl.HandleConn(c, server)
	// }))
}
