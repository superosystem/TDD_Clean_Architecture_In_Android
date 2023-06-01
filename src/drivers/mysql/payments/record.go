package payments

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/payments"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/transactions"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	ReferenceNumber string `json:"reference_number" gorm:"type:varchar(100);unique;not_null"`
	Amount          uint   `json:"amount" gorm:"not_null"`
	Status          string `json:"status" gorm:"type:enum('SUCCESS', 'PROCESS', 'PENDING', 'CANCELLED', 'ROLLBACK')" json:"status"`
	PaymentMethod   string `json:"payment_method" gorm:"type:enum('BANK', 'VA', 'MERCHANT')" json:"payment_method"`
	TransactionID   uint   `json:"transaction_id" gorm:"not_null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Transaction     *transactions.Transaction `gorm:"foreignKey:TransactionID"`
}

func FromDomain(domain *payments.Domain) *Payment {
	return &Payment{
		ID:              domain.ID,
		ReferenceNumber: domain.ReferenceNumber,
		Amount:          domain.Amount,
		Status:          domain.Status,
		PaymentMethod:   domain.PaymentMethod,
		TransactionID:   domain.TransactionID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

func (rec *Payment) ToDomain() *payments.Domain {
	return &payments.Domain{
		ID:              rec.ID,
		ReferenceNumber: rec.ReferenceNumber,
		Amount:          rec.Amount,
		Status:          rec.Status,
		PaymentMethod:   rec.PaymentMethod,
		TransactionID:   rec.TransactionID,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}
