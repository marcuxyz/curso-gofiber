package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	views := html.New("views", ".html")
	app := fiber.New(&fiber.Settings{
		Views: views,
	})

	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"Message": "Foi enviado com sucesso o email para marcus@localhost.com",
		})
	})

	app.Listen(8080)
}
