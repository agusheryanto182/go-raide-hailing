package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateImageURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	regex := `^(https?://)?([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]*)?$`
	match, _ := regexp.MatchString(regex, url)
	return match
}
