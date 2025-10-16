package main

import (
	"context"
	"hexagonal/core"
	"hexagonal/repository"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ✅ ฟังก์ชันเชื่อม MongoDB Atlas
func ConnectMongo() *mongo.Database {
	// โหลด .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("❌ MONGODB_URI not found in .env")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Cannot connect MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Ping MongoDB failed:", err)
	}

	log.Println("✅ Connected to MongoDB Atlas successfully!")

	// ใช้ database “hexagonal_db”
	return client.Database("jab")
}

func main() {
	db := ConnectMongo()
	repo := repository.NewMongoRepo(db)
	service := core.NewService(repo)

	app := fiber.New()

	// 🧠 ดึงข้อมูลทั้งหมด
	app.Get("/me", func(c *fiber.Ctx) error {
		data, err := service.GetAll()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(data)
	})

	// 🔍 ดึงข้อมูลตาม id
	app.Get("/me/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		data, err := service.GetById(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if data == nil {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.JSON(data)
	})

	// ➕ เพิ่มข้อมูลใหม่
	app.Post("/me", func(c *fiber.Ctx) error {
		var m core.Me
		if err := c.BodyParser(&m); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := service.Create(&m); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(201).JSON(m)
	})

	// ✏️ แก้ไขข้อมูล
	app.Put("/me/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var m core.Me
		if err := c.BodyParser(&m); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := service.Update(id, &m); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "updated"})
	})

	// 🗑️ ลบข้อมูล
	app.Delete("/me/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := service.Delete(id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "deleted"})
	})
	app.Get("/db", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"database": db.Name(),
			"status":   "connected",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
