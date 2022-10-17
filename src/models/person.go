package models

import (
	"github.com/duchai27798/demo_migrate/src/helpers"
	"github.com/duchai27798/demo_migrate/src/models/interfaces"
	"gorm.io/gorm"
)

type Person struct {
	interfaces.Model
	No              int8           `json:"no"`
	BirthName       string         `json:"birth_name" validate:"required"`
	GeneName        string         `json:"gene_name"`
	SecondName      string         `json:"second_name"`
	Birthday        int64          `json:"birthday"`
	DateOfDeathText string         `json:"date_of_death_text"`
	FatherId        string         `json:"father_id"`
	MotherId        string         `json:"mother_id"`
	PositionTitleId string         `json:"position_title_id" validate:"required"`
	Email           string         `json:"email"`
	PhoneNumber     string         `json:"phone_number"`
	Address         string         `json:"address"`
	Descriptions    string         `json:"descriptions"`
	StatusId        string         `json:"status_id" validate:"required"`
	Father          *Person        `json:"father,omitempty"`
	Children        *[]Person      `json:"children,omitempty" gorm:"foreignKey:FatherId"`
	PositionTitle   *PositionTitle `json:"position_title,omitempty" gorm:"foreignKey:PositionTitleId;references:Id"`
	PersonStatus    *PersonStatus  `json:"status,omitempty" gorm:"foreignKey:StatusId;references:Id"`
}

func (person *Person) Invalid() []*interfaces.ErrorResponse {
	return helpers.ValidateModel(person)
}

func (person *Person) BeforeCreate(tx *gorm.DB) (err error) {
	person.SuperBeforeCreate()
	return
}
