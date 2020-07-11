package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.All("/method", func(c *fiber.Ctx) {
		if c.Method() == "PUT" {
			c.Send("Essa mensagem foi enviada via PUT")
			return
		}

		if c.Method() == "POST" {
			c.Send("Essa mensagem foi enviada via POST")
			return
		}

		c.Send("Essa mensagem foi enviada via GET")
	})

	app.Listen(3001)
}
