package routes

import (
	"database/sql"
	"github.com/duchai27798/demo_migrate/src/controllers/api"
	"github.com/duchai27798/demo_migrate/src/controllers/auth"
	"github.com/duchai27798/demo_migrate/src/database"
	"github.com/duchai27798/demo_migrate/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DB                   *gorm.DB
	mySqlDB              *sql.DB
	constantService      services.IConstantService
	constantController   api.IConstantController
	personService        services.IPersonService
	personController     api.IPersonController
	initializeController api.IInitializeController
	authController       auth.IAuthController
)

func InitRoute(app *fiber.App) {
	DB = database.DBInstance
	mySqlDB = database.MySqlDB
	constantService = services.NewConstantService(DB)
	constantController = api.NewConstantController(constantService)
	personService = services.NewPersonService(DB)
	personController = api.NewPersonController(personService)
	initializeController = api.NewInitializeController(mySqlDB)
	authController = auth.NewAuthController()

	apiInit := app.Group("/INIT")
	{
		apiInit.Post("/DB", initializeController.InitDB)
	}

	api := app.Group("/api")
	{
		api.Get("/genders", constantController.GetGenders)
		api.Get("/position-titles", constantController.GetPositionTitles)
		api.Get("/person-statuses", constantController.GetPersonStatuses)
		api.Get("/marital-statuses", constantController.GetMaritalStatuses)

		personApi := app.Group("/person")
		{
			personApi.Get("/all", personController.GetPeople)
			personApi.Get("/get/:id", personController.GetPerson)
			personApi.Post("/create", personController.CreateNewPerson)
			personApi.Delete("/delete/:id", personController.DeletePerson)
		}
	}

	auth := app.Group("/auth")
	{
		auth.Post("/login", authController.Login)
		auth.Post("/register", authController.Register)
		auth.Post("/logout", authController.Logout)
	}
}
