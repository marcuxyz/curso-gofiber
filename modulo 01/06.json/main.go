package main

import (
	"github.com/gofiber/fiber"
)

type Person struct {
	Name string
	Age  uint8
}

func main() {
	app := fiber.New()

	app.Get("/get-name/", func(ctx *fiber.Ctx) {
		person := Person{
			Name: "Marcus Pereira",
			Age:  11,
		}
		ctx.JSON(person)
	})

	app.Listen(8080)
}
