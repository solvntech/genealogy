package models

type Person struct {
	Id              string         `json:"id"`
	BirthName       string         `json:"birth_name"`
	GeneName        string         `json:"gene_name"`
	SecondName      string         `json:"second_name"`
	Birthday        int64          `json:"birthday"`
	FatherId        string         `json:"father_id"`
	MotherId        string         `json:"mother_id"`
	PositionTitleId string         `json:"position_title_id"`
	Email           string         `json:"email"`
	PhoneNumber     string         `json:"phone_number"`
	Address         string         `json:"address"`
	Descriptions    string         `json:"descriptions"`
	StatusId        string         `json:"status_id"`
	Father          *Person        `json:"father,omitempty"`
	Children        *[]Person      `json:"children,omitempty" gorm:"foreignKey:FatherId"`
	PositionTitle   *PositionTitle `json:"position_title,omitempty" gorm:"foreignKey:PositionTitleId;references:Id"`
	PersonStatus    *PersonStatus  `json:"status,omitempty" gorm:"foreignKey:StatusId;references:Id"`
}
