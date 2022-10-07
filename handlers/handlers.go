package handlers

import (
	"UserTask/pkg"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := pkg.GetUsersDB()

	return c.JSON(users)
}

func GetUserCode(c *fiber.Ctx) error {
	login := c.Params("login")
	fmt.Println(login)

	code := pkg.GetCode(login)
	fmt.Println(code)
	return c.JSON(code)
}
