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

type UserUpdate struct {
	FullName    string `json:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Photo       string `json:"photo" form:"photo"`
}

func (req *User) ToDomainSignUp() *users.Domain {
	return &users.Domain{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}
}

func (req *SignIn) ToDomainSignIn() *users.SignInDomain {
	return &users.SignInDomain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserUpdate) ToDomainUserUpdate() users.Domain {
	return users.Domain{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
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
