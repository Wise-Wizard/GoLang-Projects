package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z] + [a-z] + [a-z] +`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) != 1
}
