package payments

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID              uint
	ReferenceNumber string
	Amount          uint
	Status          string
	PaymentMethod   string
	TransactionID   uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Repository interface {
	Create(domain *Domain) error
	Update(ID string, domain Domain) (*Domain, error)
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByReferenceNumber(referenceNumber string) (*Domain, error)
}
