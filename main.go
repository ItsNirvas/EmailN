package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("no errors found")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {

			switch v.Tag() {
			case "required":
				println(v.StructField(), "is required.")
			case "min":
				println(v.StructField(), "is required with minimum of", v.Param(), "chars.")
			case "max":
				println(v.StructField(), "is required with maximum of", v.Param(), "chars.")
			case "email":
				println(v.StructField(), "is invalid")
			}
		}
	}
}
