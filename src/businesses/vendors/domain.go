package vendors

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID          uint
	Name        string
	Price       uint
	VendorType  string
	Address     string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Repository interface {
	Create(domain *Domain) error
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByName(name string) (*Domain, error)
	GetByType(vendorType string) *[]Domain
}

type UseCase interface {
	Create(domain *Domain) error
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByName(name string) (*Domain, error)
	GetByType(vendorType string) *[]Domain
}
