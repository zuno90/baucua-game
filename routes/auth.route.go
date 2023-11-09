package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuno90/go-ws/handlers"
)

func SetUpHttpRoutes(app *fiber.App) {
	// auth
	auth := app.Group("/auth")
	auth.Post("/signup", handlers.SignUp)
	auth.Post("/login", handlers.Login)
}
