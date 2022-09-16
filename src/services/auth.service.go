package services

import (
	"github.com/duchai27798/demo_migrate/src/models/auth"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAuthService interface {
	FindUsers() *[]auth.User
	FindUser(email string) (*auth.User, error)
	CreateUser(user *auth.User) (*auth.User, error)
	DeleteUser(id string) (*auth.User, error)
}

type AuthService struct {
	DB *gorm.DB
}

func (authService AuthService) FindUsers() *[]auth.User {
	var users *[]auth.User
	authService.DB.Model(&auth.User{}).Preload("Role").Find(&users)
	return users
}

func (authService AuthService) FindUser(email string) (*auth.User, error) {
	var user *auth.User
	if db := authService.DB.Where("email = ?", email).Preload(clause.Associations).Preload("Role").First(&user); db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}

func (authService AuthService) CreateUser(user *auth.User) (*auth.User, error) {
	tx := authService.DB.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (authService AuthService) DeleteUser(id string) (*auth.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(DB *gorm.DB) IAuthService {
	return &AuthService{
		DB,
	}
}
