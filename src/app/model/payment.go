package model

import "time"

type Payment struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	ReferenceNumber string `gorm:"type:varchar(100)" json:"reference_number"`
	Amount          uint   `gorm:"type:uint;not_null" json:"amount"`
	Status          string `gorm:"type:enum('SUCCESS', 'PENDING', 'CANCELLED');not_null" json:"status"`
	PaymentMethod   string `gorm:"type:varchar(100)" json:"payment_method"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TransactionID   uint         `json:"transaction_id"`
	Transaction     *Transaction `gorm:"foreignKey:TransactionID"`
}
