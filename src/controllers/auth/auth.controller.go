package auth

import (
	"github.com/duchai27798/demo_migrate/src/helpers"
	"github.com/duchai27798/demo_migrate/src/models/auth"
	"github.com/duchai27798/demo_migrate/src/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type IAuthenticationController interface {
	Login(context *fiber.Ctx) error
	GetAllUsers(context *fiber.Ctx) error
	Register(context *fiber.Ctx) error
	Logout(context *fiber.Ctx) error
}

type AuthenticationController struct {
	authService services.IAuthService
	jwtService  services.IJWTService
}

func (authenticationController AuthenticationController) GetAllUsers(context *fiber.Ctx) error {
	users := authenticationController.authService.FindUsers()
	return context.Status(fiber.StatusOK).JSON(users)
}

func (authenticationController AuthenticationController) Login(context *fiber.Ctx) error {
	userLogin := &auth.UserLogin{}
	context.BodyParser(userLogin)

	user, err := authenticationController.authService.FindUser(userLogin.Email)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(helpers.Response(fiber.StatusNotFound, "User not found"))
	}

	loginErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))

	if loginErr != nil {
		return context.Status(fiber.StatusBadRequest).JSON(helpers.Response(fiber.StatusBadRequest, "Incorrect password"))
	}

	// credential
	claim := &auth.JWTClaim{
		UserId: user.ID,
		Email:  user.Email,
		RoleId: user.RoleId,
	}

	tokenString, tokenErr := authenticationController.jwtService.GenericToken(claim)

	if tokenErr != nil {
		return context.Status(fiber.StatusBadRequest).JSON(helpers.Response(fiber.StatusBadRequest, tokenErr.Error()))
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"user_id": user.ID,
		"email":   user.Email,
		"role_id": user.RoleId,
		"token":   tokenString,
	})
}

func (authenticationController AuthenticationController) Register(context *fiber.Ctx) error {
	user := &auth.User{}
	context.BodyParser(user)

	if err := user.Invalid(); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(hashPassword)

	user, errCreated := authenticationController.authService.CreateUser(user)

	if errCreated != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errCreated)
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func (authenticationController AuthenticationController) Logout(context *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthController(authService services.IAuthService, jwtService services.IJWTService) IAuthenticationController {
	return &AuthenticationController{
		authService,
		jwtService,
	}
}
