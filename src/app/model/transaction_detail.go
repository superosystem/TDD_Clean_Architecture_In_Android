package model

import "time"

type TransactionDetail struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ReferenceNumber string    `gorm:"type:varchar(100)" json:"reference_number"`
	Name            string    `gorm:"type:varchar(100)" json:"name"`
	EventDate       time.Time `json:"event_date"`
	Venue           string    `gorm:"type:varchar(100)" json:"venue"`
	Decoration      string    `gorm:"type:varchar(100)" json:"decoration"`
	Catering        string    `gorm:"type:varchar(100)" json:"catering"`
	Mua             string    `gorm:"type:varchar(100)" json:"mua"`
	Documentary     string    `gorm:"type:varchar(100)" json:"documentary"`
}
