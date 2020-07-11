package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Put("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Delete("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Patch("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Listen(8080)
}
