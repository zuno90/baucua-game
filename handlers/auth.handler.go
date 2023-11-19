package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	en "github.com/zuno90/go-ws/entities"
	st "github.com/zuno90/go-ws/handlers/game"
	"github.com/zuno90/go-ws/utils"
)

func SignUp(c *fiber.Ctx) error {
	fmt.Println("sign up!")
	return c.Status(fiber.StatusOK).JSON(st.Resp(true, "sign up!", nil, nil))
}

func Login(c *fiber.Ctx) error {
	userInput := en.UserInput{}
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(st.Resp(false, "Bad input!", nil, nil))
	}

	fmt.Printf("username: %s password: %s", userInput.Username, userInput.Password)
	// validate input
	if errors := utils.ValidateStruct(userInput); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(st.Resp(false, "", nil, errors))
	}
	// check database

	// response jwt
	return c.Status(fiber.StatusOK).JSON(st.Resp(true, "sign in!", nil, nil))
}
