package schedules

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/schedules"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/vendors"
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	ID              uint            `json:"id" gorm:"primaryKey"`
	ReferenceNumber string          `json:"reference_number" gorm:"type:varchar(100);unique;not_null"`
	Description     string          `json:"description" gorm:"type:varchar(255)"`
	Status          string          `json:"status" gorm:"type:enum('ONGOING', 'DONE', 'CANCELLED')" json:"vendor_type"`
	Date            string          `json:"date" gorm:"type:varchar(50);not_null"`
	Time            string          `json:"time" gorm:"type:varchar(50);not_null"`
	Place           string          `json:"place" gorm:"type:varchar(255);not_null"`
	UserID          uint            `json:"user_id" gorm:"not_null"`
	VendorID        uint            `json:"vendor_id" gorm:"not_null"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `json:"deleted_at"`
	User            *users.User     `gorm:"foreignKey:UserID"`
	Vendor          *vendors.Vendor `gorm:"foreignKey:VendorID"`
}

func FromDomain(domain *schedules.Domain) *Schedule {
	return &Schedule{
		ID:              domain.ID,
		ReferenceNumber: domain.ReferenceNumber,
		Description:     domain.Description,
		Status:          domain.Status,
		Date:            domain.Date,
		Time:            domain.Time,
		Place:           domain.Place,
		UserID:          domain.UserID,
		VendorID:        domain.VendorID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

func (rec *Schedule) ToDomain() *schedules.Domain {
	return &schedules.Domain{
		ID:              rec.ID,
		ReferenceNumber: rec.ReferenceNumber,
		Description:     rec.Description,
		Status:          rec.Status,
		Date:            rec.Date,
		Time:            rec.Time,
		Place:           rec.Place,
		UserID:          rec.UserID,
		VendorID:        rec.VendorID,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}
