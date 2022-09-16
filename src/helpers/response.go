package helpers

import "github.com/gofiber/fiber/v2"

func Response(status int, message interface{}) *fiber.Map {
	return &fiber.Map{
		"status":  status,
		"message": message,
	}
}
