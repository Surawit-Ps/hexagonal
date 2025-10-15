package main

import (
	"context"
	"hexagonal/core"
	"hexagonal/repository"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	// Connect MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("hexagonal_db")

	// Initialize repository + service
	repo := repository.NewMongoRepo(db)
	service := core.NewService(repo)

	// Routes
	app.Get("/me", func(c *fiber.Ctx) error {
		data, err := service.GetAll()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(data)
	})

	app.Post("/me", func(c *fiber.Ctx) error {
		var m core.Me
		if err := c.BodyParser(&m); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		if err := service.Create(&m); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(m)
	})

	log.Fatal(app.Listen(":3000"))
}
