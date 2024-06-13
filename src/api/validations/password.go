package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/nazarow/golang-clean/common"
)

func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
