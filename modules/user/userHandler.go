package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func userRegister(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.SendString(err.Error())
	}
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are missing some fields please check again"})
	}
	results, err := _userRegister(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": results})
}

func userLogin(c *fiber.Ctx) error {
	userRequest := new(User)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if userRequest.Email == "" || userRequest.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are missing some fields please check again"})
	}

	fmt.Println(userRequest)
	results, err := _userLogin(*userRequest)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"token": results})
}
