package auth

import (
	"github.com/duchai27798/demo_migrate/src/helpers"
	"github.com/duchai27798/demo_migrate/src/models"
	"github.com/duchai27798/demo_migrate/src/models/interfaces"
	"gorm.io/gorm"
)

type User struct {
	interfaces.Model
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required"`
	RoleId   string      `json:"role_id"`
	Role     models.Role `json:"role,omitempty" gorm:"foreignKey:RoleId;references:Id"`
}

func (user *User) Invalid() []*interfaces.ErrorResponse {
	return helpers.ValidateModel(user)
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.SuperBeforeCreate()
	return
}
