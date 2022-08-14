package helpers

import (
	"github.com/duchai27798/demo_migrate/src/models/interfaces"
	"github.com/duchai27798/demo_migrate/src/validators"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func InitValidator() {
	validate.RegisterValidation("timestamp", validators.TimestampValidator)
}

func ValidateModel(model interfaces.IModel) []*interfaces.ErrorResponse {
	var errors []*interfaces.ErrorResponse
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element interfaces.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
