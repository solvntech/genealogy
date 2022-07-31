package services

import (
	"fmt"
	"github.com/duchai27798/demo_migrate/src/models"
	"gorm.io/gorm"
)

type IConstantService interface {
	FindGenders() *[]models.Gender
	FindPositionTitles() *[]models.PositionTitle
	FindPersonStatuses() *[]models.PersonStatus
	FindMaritalStatuses() *[]models.MaritalStatus
}

type ConstantService struct {
	DB *gorm.DB
}

func (constantService ConstantService) FindPersonStatuses() *[]models.PersonStatus {
	var personStatuses *[]models.PersonStatus
	constantService.DB.Find(&personStatuses)
	return personStatuses
}

func (constantService ConstantService) FindMaritalStatuses() *[]models.MaritalStatus {
	var maritalStatuses *[]models.MaritalStatus
	constantService.DB.Find(&maritalStatuses)
	return maritalStatuses
}

func (constantService ConstantService) FindGenders() *[]models.Gender {
	var genders *[]models.Gender
	constantService.DB.Find(&genders)
	return genders
}

func (constantService ConstantService) FindPositionTitles() *[]models.PositionTitle {
	var positionTitles *[]models.PositionTitle
	constantService.DB.Model(&models.PositionTitle{}).Preload("Gender").Find(&positionTitles)
	return positionTitles
}

func NewConstantService(DB *gorm.DB) IConstantService {
	fmt.Println(DB)
	return &ConstantService{
		DB,
	}
}
