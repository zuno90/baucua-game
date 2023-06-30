package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/zuno90/go-ws/handlers"
	"github.com/zuno90/go-ws/structs"
)

func SetUpWebsocket(app *fiber.App) {
	pool := structs.NewPool()
	go pool.Start()
	// websocket
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Query("jwt")) // 1.0
	}))

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// log.Println(c.Locals("allowed"))  // true
		// log.Println(c.Params("id"))       // 123
		// log.Println(c.Query("v"))         // 1.0
		// log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		handlers.HandleConn(pool, c)
	}))
}
