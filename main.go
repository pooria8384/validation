package main

import (
	"new-proj/models"
	"new-proj/validators"
)

func main() {
	user := models.User{
		Username: "pooria",
		Email:    "pooria@gmail.com",
		Age:      19,
		Phone:    "09312345678",
	}

	validators.ValidatStruct(user)

}
