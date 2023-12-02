package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"password" xml:"password" form:"password"`
}

func Server() *fiber.App {
	rc := initDb()

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

		err := rc.SignUp(user.Name, user.Pass, c.Context())

		if err != nil {
			fmt.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"status": "error", "error": "Something went wrong"})
		}

		rc.GetAllUsers()

		return c.SendString("Signup Success")
	})

	return app
}
