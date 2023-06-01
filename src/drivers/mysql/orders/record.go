package orders

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	ReferenceNumber  string         `json:"reference_number" gorm:"type:varchar(100);unique;not_null"`
	TotalAmount      uint           `json:"total_amount" gorm:"not_null"`
	Venue            string         `json:"venue" gorm:"type:varchar(100)"`
	VenuePrice       uint           `json:"venue_price" gorm:"not_null"`
	Decoration       string         `json:"decoration" gorm:"type:varchar(100)"`
	DecorationPrice  uint           `json:"decoration_price" gorm:"not_null"`
	Catering         string         `json:"catering" gorm:"type:varchar(100)"`
	CateringPrice    uint           `json:"catering_price" gorm:"not_null"`
	Mua              string         `json:"mua" gorm:"type:varchar(100)"`
	MuaPrice         uint           `json:"mua_price" gorm:"not_null"`
	Documentary      string         `json:"documentary" gorm:"type:varchar(100)"`
	DocumentaryPrice uint           `json:"documentary_price" gorm:"not_null"`
	UserID           uint           `json:"user_id" gorm:"not_null"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	User             *users.User    `gorm:"foreignKey:UserID"`
}

func FromDomain(domain *orders.Domain) *Order {
	return &Order{
		ID:               domain.ID,
		ReferenceNumber:  domain.ReferenceNumber,
		TotalAmount:      domain.TotalAmount,
		Venue:            domain.Venue,
		VenuePrice:       domain.VenuePrice,
		Decoration:       domain.Decoration,
		DecorationPrice:  domain.DecorationPrice,
		Catering:         domain.Catering,
		CateringPrice:    domain.CateringPrice,
		Mua:              domain.Mua,
		MuaPrice:         domain.MuaPrice,
		Documentary:      domain.Documentary,
		DocumentaryPrice: domain.DocumentaryPrice,
		UserID:           domain.UserID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
	}
}

func (rec *Order) ToDomain() *orders.Domain {
	return &orders.Domain{
		ID:              rec.ID,
		ReferenceNumber: rec.ReferenceNumber,
		TotalAmount:     rec.TotalAmount,
		Venue:           rec.Venue,
		Decoration:      rec.Decoration,
		Catering:        rec.Catering,
		Mua:             rec.Mua,
		Documentary:     rec.Documentary,
		UserID:          rec.UserID,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}
