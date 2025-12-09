package validator

import (
	"regexp"

	validatorLib "github.com/go-playground/validator/v10"
)

var (
	hasUpper   = regexp.MustCompile(`[A-Z]`)
	hasDigit   = regexp.MustCompile(`[0-9]`)
	minLength8 = regexp.MustCompile(`.{8,}`)
)

func passwordStrong(field validatorLib.FieldLevel) bool {
	s := field.Field().String()
	return hasUpper.MatchString(s) && hasDigit.MatchString(s) && minLength8.MatchString(s)
}

func SetupValidator() (*validatorLib.Validate, error) {
	v := validatorLib.New()
	err := v.RegisterValidation("password_strong", passwordStrong)

	return v, err
}
