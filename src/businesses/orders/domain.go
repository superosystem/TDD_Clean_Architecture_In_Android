package orders

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID               uint
	ReferenceNumber  string
	Status           string
	TotalAmount      uint
	Venue            string
	VenuePrice       uint
	Decoration       string
	DecorationPrice  uint
	Catering         string
	CateringPrice    uint
	Mua              string
	MuaPrice         uint
	Documentary      string
	DocumentaryPrice uint
	UserID           uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type Repository interface {
	Create(domain *Domain) (*Domain, error)
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByReferenceNumber(referenceNumber string) (*Domain, error)
}

type UseCase interface {
	Create(domain *Domain) (*Domain, error)
	Update(ID string, domain Domain) (*Domain, error)
	Delete(ID string) bool
	GetAll() *[]Domain
	GetByID(ID string) (*Domain, error)
	GetByReferenceNumber(referenceNumber string) (*Domain, error)
}
