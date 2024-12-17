package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator"
)

func main() {
	contacts := []campaign.Contact{{Email: ""}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is less than the minimum")
			case "max":
				println(v.StructField() + " is greater than the maximum")
			case "email":
				println(v.StructField() + " is not a valid email")
			}

			println(v.StructField() + " is invalid:" + v.Tag())
		}
	}
}
