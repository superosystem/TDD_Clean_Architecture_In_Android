package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
)

type User struct {
	FullName             string `json:"full_name" validate:"required"`
	Email                string `json:"email" validate:"required,email,lowercase"`
	Password             string `json:"password" validate:"required"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
	PhoneNumber          string `json:"phone_number"`
	Photo                string `json:"photo" form:"photo"`
	Roles                string `json:"roles"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required"`
}

func (req *User) ToDomainSignUp() *users.Domain {
	return &users.Domain{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *SignIn) ToDomainSignIn() *users.SignInDomain {
	return &users.SignInDomain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *SignIn) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
