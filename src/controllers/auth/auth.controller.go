package auth

import "github.com/gofiber/fiber/v2"

type IAuthController interface {
	Login(context *fiber.Ctx) error
	Register(context *fiber.Ctx) error
	Logout(context *fiber.Ctx) error
}

type AuthController struct {
}

func (authController AuthController) Login(context *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (authController AuthController) Register(context *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (authController AuthController) Logout(context *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthController() IAuthController {
	return &AuthController{}
}
