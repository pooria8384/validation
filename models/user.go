package models

type User struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=18,lte=65"`
	Phone    string `validate:"required,phone_Number"`
}
