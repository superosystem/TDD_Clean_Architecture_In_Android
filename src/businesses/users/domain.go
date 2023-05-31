package users

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID          uint
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Photo       string
	Roles       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type SignUpDomain struct {
	FullName string
	Email    string
	Password string
}

type SignInDomain struct {
	Email    string
	Password string
}

type TokenDomain struct {
	Token string
}

type Repository interface {
	Create(domain *Domain) error
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
}

type UseCase interface {
	SignUp(domain *Domain) error
	SignIn(domain *SignInDomain) (interface{}, error)
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
}
