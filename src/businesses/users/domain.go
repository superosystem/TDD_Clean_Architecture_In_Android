package users

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID        uint
	FullName  string
	Email     string
	Password  string
	Photo     string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type SignUpDomain struct {
	FullName string
	Email    string
	Password string
}
type SignInDomain struct {
	FullName string
	Email    string
	Password string
}

type Repository interface {
	Create(domain *Domain) error
	Update(ID string, domain *Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() []Domain
	GetByID(ID string) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
	IsExistUser(ID, email string) bool
}

type UseCase interface {
	SignUp(domain *Domain) error
	SignIn(domain *SignInDomain) (*Domain, error)
	Update(ID string, domain *Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() []Domain
	GetByID(ID string) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
}
