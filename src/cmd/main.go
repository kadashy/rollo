package main

import (
	v1 "rollo/internal/server/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//port := os.Getenv("PORT")
	port := ":8082"
	app := fiber.New()
	app.Use(cors.New())

	v1.SetupRoutes(app)
	app.Listen(port)

}
