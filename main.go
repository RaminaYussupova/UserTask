package main

import (
	"UserTask/handlers"
	"UserTask/pkg"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	Init()
	app := fiber.New()

	app.Get("/getusers", handlers.GetUsers)
	app.Get("/getusercode/:login", handlers.GetUserCode)

	app.Listen(":3000")

}
func Init() {
	pkg.InitDB()
}
