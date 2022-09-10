package services

import (
	"github.com/duchai27798/demo_migrate/src/models"
	"gorm.io/gorm"
)

type IAuthService interface {
	FindUsers() *[]models.User
	FindUser(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) (*models.User, error)
}

type AuthService struct {
	DB *gorm.DB
}

func (authService AuthService) FindUsers() *[]models.User {
	var users *[]models.User
	authService.DB.Model(&models.User{}).Preload("Role").Find(&users)
	return users
}

func (authService AuthService) FindUser(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (authService AuthService) CreateUser(user *models.User) (*models.User, error) {
	tx := authService.DB.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (authService AuthService) DeleteUser(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(DB *gorm.DB) IAuthService {
	return &AuthService{
		DB,
	}
}
