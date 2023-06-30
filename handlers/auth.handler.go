package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zuno90/go-ws/entities"
	st "github.com/zuno90/go-ws/structs"
)

func SignUp(c *fiber.Ctx) error {
	fmt.Println("sign up!")
	return c.Status(http.StatusOK).JSON(st.Resp(true, "sign up!", nil))
}

func Login(c *fiber.Ctx) error {
	user := entities.UserInput{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(st.Resp(false, "Bad input!", nil))
	}

	fmt.Printf("username: %s password: %s", user.Username, user.Password)
	// validate input

	// check database

	// response jwt
	return c.Status(http.StatusOK).JSON(st.Resp(true, "sign in!", nil))
}
