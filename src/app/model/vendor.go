package model

import "time"

type Vendor struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(100);not_null;unique" json:"name"`
	Price       uint    `gorm:"type:uint;not_null" json:"price"`
	Type        string  `gorm:"type:enum('VENUE', 'DECORATION', 'CATERING', 'DOCUMENTARY');not_null" json:"role"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Image       *string `gorm:"type:longtext" json:"image"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
