package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuno90/go-ws/handlers"
)

func SetUpRoutes(app *fiber.App) {
	// auth
	app.Post("/auth/signup", handlers.SignUp)
	app.Post("/auth/login", handlers.Login)
}
