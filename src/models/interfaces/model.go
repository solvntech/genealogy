package interfaces

import (
	"github.com/google/uuid"
	"time"
)

type IModel interface {
	Invalid() []*ErrorResponse
}

type Model struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

//func (model *Model) BeforeCreate(tx *gorm.DB) (err error) {
//	model.SuperBeforeCreate()
//	return
//}

func (model *Model) SuperBeforeCreate() {
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now().Unix()
	model.UpdatedAt = time.Now().Unix()
}
