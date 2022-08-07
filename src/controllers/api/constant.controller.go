package api

import (
	"github.com/duchai27798/demo_migrate/src/services"
	"github.com/gofiber/fiber/v2"
)

type IConstantController interface {
	GetGenders(context *fiber.Ctx) error
	GetPositionTitles(context *fiber.Ctx) error
	GetMaritalStatuses(context *fiber.Ctx) error
	GetPersonStatuses(context *fiber.Ctx) error
}

type ConstantController struct {
	constantService services.IConstantService
}

func (constantController ConstantController) GetMaritalStatuses(context *fiber.Ctx) error {
	personStatuses := constantController.constantService.FindMaritalStatuses()
	return context.Status(fiber.StatusOK).JSON(personStatuses)
}

func (constantController ConstantController) GetPersonStatuses(context *fiber.Ctx) error {
	personStatuses := constantController.constantService.FindPersonStatuses()
	return context.Status(fiber.StatusOK).JSON(personStatuses)
}

func (constantController ConstantController) GetGenders(context *fiber.Ctx) error {
	genders := constantController.constantService.FindGenders()
	return context.Status(fiber.StatusOK).JSON(genders)
}

func (constantController ConstantController) GetPositionTitles(context *fiber.Ctx) error {
	positionTitles := constantController.constantService.FindPositionTitles()
	return context.Status(fiber.StatusOK).JSON(positionTitles)
}

func NewConstantController(constantService services.IConstantService) IConstantController {
	return &ConstantController{
		constantService,
	}
}
