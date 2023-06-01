package schedules

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID              uint
	ReferenceNumber string
	Description     string
	Status          string
	Date            string
	Time            string
	Place           string
	UserID          uint
	VendorID        uint
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
