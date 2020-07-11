package main

import "github.com/gofiber/fiber"

type Email struct {
	To string
}

func main() {
	app := fiber.New()

	app.Post("/email", func(c *fiber.Ctx) {
		var email Email
		c.BodyParser(&email)
		c.Status(201).Send("Mensagem enviada com sucesso para: " + email.To)
	})

	app.Listen(3001)
}
