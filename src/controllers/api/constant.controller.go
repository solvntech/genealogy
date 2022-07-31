package api

import (
	"github.com/duchai27798/demo_migrate/src/services"
	"github.com/gofiber/fiber/v2"
)

type IConstantController interface {
	GetGenders(content *fiber.Ctx) error
	GetPositionTitles(content *fiber.Ctx) error
	GetMaritalStatuses(content *fiber.Ctx) error
	GetPersonStatuses(content *fiber.Ctx) error
}

type ConstantController struct {
	constantService services.IConstantService
}

func (constantController ConstantController) GetMaritalStatuses(content *fiber.Ctx) error {
	personStatuses := constantController.constantService.FindMaritalStatuses()
	return content.Status(fiber.StatusOK).JSON(personStatuses)
}

func (constantController ConstantController) GetPersonStatuses(content *fiber.Ctx) error {
	personStatuses := constantController.constantService.FindPersonStatuses()
	return content.Status(fiber.StatusOK).JSON(personStatuses)
}

func (constantController ConstantController) GetGenders(content *fiber.Ctx) error {
	genders := constantController.constantService.FindGenders()
	return content.Status(fiber.StatusOK).JSON(genders)
}

func (constantController ConstantController) GetPositionTitles(content *fiber.Ctx) error {
	positionTitles := constantController.constantService.FindPositionTitles()
	return content.Status(fiber.StatusOK).JSON(positionTitles)
}

func NewConstantController(constantService services.IConstantService) IConstantController {
	return &ConstantController{
		constantService,
	}
}
