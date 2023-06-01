package vendors

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"gorm.io/gorm"
	"time"
)

type Vendor struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100);unique;not_null"`
	Price       uint           `json:"price" gorm:"not_null"`
	VendorType  string         `json:"vendor_type" gorm:"type:enum('VENUE', 'DECORATION', 'CATERING', 'MUA', 'DOCUMENTARY')" json:"vendor_type"`
	Address     string         `json:"address" gorm:"type:varchar(255)"`
	Description string         `json:"description" gorm:"type:varchar(255)"`
	Image       string         `json:"image" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *vendors.Domain) *Vendor {
	return &Vendor{
		ID:          domain.ID,
		Name:        domain.Name,
		Price:       domain.Price,
		VendorType:  domain.VendorType,
		Description: domain.Description,
		Image:       domain.Image,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (rec *Vendor) ToDomain() *vendors.Domain {
	return &vendors.Domain{
		ID:          rec.ID,
		Name:        rec.Name,
		Price:       rec.Price,
		VendorType:  rec.VendorType,
		Description: rec.Description,
		Image:       rec.Image,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}
