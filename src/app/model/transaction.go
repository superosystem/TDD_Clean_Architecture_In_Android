package model

import "time"

type Transaction struct {
	ID                  uint   `gorm:"primaryKey" json:"id"`
	ReferenceNumber     string `gorm:"type:varchar(100)" json:"reference_number"`
	TotalAmount         uint   `gorm:"type:uint;not_null" json:"total"`
	Payment             uint   `gorm:"type:uint;not_null" json:"paid"`
	Debt                uint   `gorm:"type:uint;not_null" json:"debt"`
	Status              string `gorm:"type:enum('SUCCESS', 'PENDING', 'CANCELLED');not_null" json:"status"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              uint               `json:"user_id"`
	TransactionDetailID uint               `json:"transaction_detail_id"`
	User                *User              `gorm:"foreignKey:UserID"`
	TransactionDetail   *TransactionDetail `gorm:"foreignKey:TransactionDetailID"`
}
