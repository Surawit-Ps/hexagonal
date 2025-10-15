package main

import "github.com/gofiber/fiber/v2"

func main() {
	print("hello")
	app := fiber.New( )

	app.Get("/kub",func(c *fiber.Ctx) error{
		return c.SendString("momo")
	})
}