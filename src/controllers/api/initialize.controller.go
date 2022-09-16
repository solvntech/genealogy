package api

import (
	"database/sql"
	"github.com/duchai27798/demo_migrate/src/database"
	"github.com/duchai27798/demo_migrate/src/helpers"
	"github.com/gofiber/fiber/v2"
)

type IInitializeController interface {
	InitDB(content *fiber.Ctx) error
}

type InitializeController struct {
	mysqlSql *sql.DB
}

func (initializeController InitializeController) InitDB(context *fiber.Ctx) error {
	if err := database.Migration(initializeController.mysqlSql); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helpers.Response(fiber.StatusInternalServerError, err.Error()))
	}
	return context.Status(fiber.StatusOK).JSON(helpers.Response(fiber.StatusOK, "Init DB successfully"))
}

func NewInitializeController(mysqlSql *sql.DB) IInitializeController {
	return &InitializeController{
		mysqlSql,
	}
}
