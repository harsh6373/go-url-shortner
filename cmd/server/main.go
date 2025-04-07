package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harsh6373/go-url-shortner/config"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Define routes (to be implemented)
	// app.Get("/", handler.Function)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
