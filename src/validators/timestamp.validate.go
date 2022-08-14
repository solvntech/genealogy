package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func TimestampValidator(fl validator.FieldLevel) bool {
	timestampRegexpString := "^([0-9]{10}|[0-9]{13})$"
	value := fmt.Sprintf("%v", fl.Field().Interface())

	ok, _ := regexp.MatchString(timestampRegexpString, value)
	fmt.Println(value)
	return ok
}
