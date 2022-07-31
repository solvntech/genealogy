package main

import (
	"fmt"
	"github.com/duchai27798/demo_migrate/src/database"
	"github.com/duchai27798/demo_migrate/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	database.ConnectDB()

	app := fiber.New()
	routes.InitRoute(app)
	app.Listen(":3000")
}
