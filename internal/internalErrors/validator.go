package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "min":
		return errors.New(field + " is required with a minimum of " + validationError.Param())
	case "max":
		return errors.New(field + " is required with a maximum of " + validationError.Param())
	case "email":
		return errors.New(field + " is required with a valid email")
	}
	return nil
}
