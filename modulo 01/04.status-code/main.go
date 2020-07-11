package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("GET")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		if "ok" != "ok1" {
			ctx.Status(http.StatusNotFound).Send("A página não foi encontrada!")
			return
		}

		ctx.Status(http.StatusCreated).Send("POST")
	})

	app.Listen(8080)
}
