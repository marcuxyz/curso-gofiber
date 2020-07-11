package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("GET")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Send("POST")
	})

	app.Put("/", func(ctx *fiber.Ctx) {
		ctx.Send("PUT")
	})

	app.Delete("/", func(ctx *fiber.Ctx) {
		ctx.Send("DELETE")
	})

	app.Patch("/", func(ctx *fiber.Ctx) {
		ctx.Send("PATCH")
	})

	app.Listen(8080)
}
