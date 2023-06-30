package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zuno90/go-ws/configs"
	"github.com/zuno90/go-ws/routes"
)

func main() {
	log.SetFlags(log.Lshortfile)

	env := os.Getenv("GO_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
	// connect database
	configs.ConnectPostgresDB()

	// init httpserver/websocket
	app := fiber.New()

	routes.SetUpRoutes(app)
	routes.SetUpWebsocket(app)

	log.Fatal(app.Listen(":3000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
