package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"password" xml:"password" form:"password"`
}

func Server() *fiber.App {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1 := app.Group("/api/v1")

	v1.Post("/signup", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			log.Fatal(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"status": "error", "error": err})
		}
		return c.SendString("Signup")
	})

	return app

}
