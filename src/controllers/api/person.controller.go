package api

import (
	"github.com/duchai27798/demo_migrate/src/models"
	"github.com/duchai27798/demo_migrate/src/services"
	"github.com/gofiber/fiber/v2"
)

type IPersonController interface {
	GetPeople(context *fiber.Ctx) error
	GetPerson(context *fiber.Ctx) error
	CreateNewPerson(context *fiber.Ctx) error
	DeletePerson(context *fiber.Ctx) error
}

type PersonController struct {
	personService services.IPersonService
}

func (personController PersonController) GetPerson(context *fiber.Ctx) error {
	personId := context.Params("id")
	person, err := personController.personService.FindPerson(personId)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return context.Status(fiber.StatusOK).JSON(person)
}

func (personController PersonController) DeletePerson(context *fiber.Ctx) error {
	personId := context.Params("id")
	person, err := personController.personService.DeletePerson(personId)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return context.Status(fiber.StatusOK).JSON(person)
}

func (personController PersonController) GetPeople(context *fiber.Ctx) error {
	people := personController.personService.FindPeople()
	return context.Status(fiber.StatusOK).JSON(people)
}

func (personController PersonController) CreateNewPerson(context *fiber.Ctx) error {
	person := &models.Person{}
	context.BodyParser(person)

	if err := person.Invalid(); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	person, errCreated := personController.personService.CreatePerson(person)

	if errCreated != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errCreated)
	}

	return context.Status(fiber.StatusOK).JSON(person)
}

func NewPersonController(personService services.IPersonService) IPersonController {
	return &PersonController{
		personService,
	}
}
