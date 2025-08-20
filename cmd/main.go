package main

import (
	"log"
	"os"

	"inventory-system/config"
	"inventory-system/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	config.ConnectRedis()

	// Create Fiber app
	app := fiber.New()

	// Register routes
	routes.SetupRoutes(app)

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
