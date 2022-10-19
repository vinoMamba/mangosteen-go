package requests

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type EmailExistRequest struct {
	Email string `json:"email" validate:"email,required"`
}

func ValidateEmailExistRequest(request EmailExistRequest) bool {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
