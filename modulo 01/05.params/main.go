package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/get-name/:name", func(ctx *fiber.Ctx) {
		name := ctx.Params("name")
		ctx.Send("Meu nome é: " + name)
	})

	app.Listen(8080)
}
