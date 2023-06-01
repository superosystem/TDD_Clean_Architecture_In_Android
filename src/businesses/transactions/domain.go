package transactions

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID              uint
	ReferenceNumber string
	Status          string
	Cred            uint
	Debt            uint
	Total           uint
	OrderID         uint
	UserID          uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Repository interface {
	Create(domain *Domain) (*Domain, error)
	Update(ID string, domain Domain) (*Domain, error)
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByReferenceNumber(referenceNumber string) (*Domain, error)
}

type UseCase interface {
	Create(domain *Domain) (*Domain, error)
	GetAll() *[]Domain
}
