package services

import (
	"github.com/duchai27798/demo_migrate/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPersonService interface {
	FindPeople() *[]models.Person
	FindPerson(id string) (*models.Person, error)
	CreatePerson(id string) (*models.Person, error)
}

type PersonService struct {
	DB *gorm.DB
}

func (personService PersonService) CreatePerson(id string) (*models.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personService PersonService) FindPerson(id string) (*models.Person, error) {
	var person *models.Person
	if db := personService.DB.Where("id = ?", id).Preload(clause.Associations).Preload("Children.PositionTitle").Preload("Children.PersonStatus").First(&person); db.Error != nil {
		return nil, db.Error
	}
	return person, nil
}

func (personService PersonService) FindPeople() *[]models.Person {
	var people *[]models.Person
	personService.DB.Preload("PositionTitle").Preload("PositionTitle.Gender").Preload("PersonStatus").Find(&people)
	return people
}

func NewPersonService(DB *gorm.DB) IPersonService {
	return &PersonService{
		DB,
	}
}
