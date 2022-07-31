package models

type PositionTitle struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	GenderId string `json:"gender_id"`
	Gender   Gender `json:"gender,omitempty" gorm:"foreignKey:GenderId;references:Id"`
}
