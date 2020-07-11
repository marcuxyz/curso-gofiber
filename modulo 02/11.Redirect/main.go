package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/redirect", func(c *fiber.Ctx) {
		c.Redirect("/home")
	})

	app.Get("/home", func(c *fiber.Ctx) {
		c.Send("Você está vindo de uma rota de redirecionamento. /redirect")
	})

	app.Listen(8080)
}
