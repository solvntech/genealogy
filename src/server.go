package main

import (
	"fmt"
	"github.com/duchai27798/demo_migrate/src/database"
	"github.com/duchai27798/demo_migrate/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	database.ConnectDB()
	ApiPort := os.Getenv("API_PORT")
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	routes.InitRoute(app)
	app.Listen(":" + ApiPort)
}
