package main

import (
	v1 "rollo/internal/server/v1"

	"github.com/gofiber/fiber"
)

func main() {
	//port := os.Getenv("PORT")
	port := 8082
	app := fiber.New()

	v1.SetupRoutes(app)
	app.Listen(port)

}
