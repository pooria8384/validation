package validators

import (
	"fmt"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

func PhoneValidation(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^(\+98|0)?(9(1[0-9]|3[0-9]|2[0-9]|0[0-9]|9[0-2]))\d{7}$`)

	phone, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return regex.MatchString(phone)
}

func ValidatStruct(input interface{}) {
	validate := validator.New()

	validate.RegisterValidation("phone_Number", PhoneValidation)

	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Validation failed for field '%s': rule '%s' \n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("validation was succsusful")
	}
}
