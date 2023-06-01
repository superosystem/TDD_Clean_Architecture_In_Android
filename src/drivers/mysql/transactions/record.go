package transactions

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/transactions"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/orders"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	ReferenceNumber string         `json:"reference_number" gorm:"type:varchar(100);unique;not_null"`
	Status          string         `json:"status" gorm:"type:enum('PAID OFF', 'PROCESS', 'CANCELLED')" json:"status"`
	Cred            uint           `json:"cred" gorm:"not_null"`
	Debt            uint           `json:"debt" gorm:"not_null"`
	Total           uint           `json:"total" gorm:"not_null"`
	OrderID         uint           `json:"order_id" gorm:"not_null"`
	UserID          uint           `json:"user_id" gorm:"not_null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	Order           *orders.Order  `gorm:"foreignKey:OrderID"`
	User            *users.User    `gorm:"foreignKey:UserID"`
}

func FromDomain(domain *transactions.Domain) *Transaction {
	return &Transaction{
		ID:              domain.ID,
		ReferenceNumber: domain.ReferenceNumber,
		Status:          domain.Status,
		Cred:            domain.Cred,
		Debt:            domain.Debt,
		Total:           domain.Total,
		OrderID:         domain.OrderID,
		UserID:          domain.UserID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

func (rec *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		ID:              rec.ID,
		ReferenceNumber: rec.ReferenceNumber,
		Status:          rec.Status,
		Cred:            rec.Cred,
		Debt:            rec.Debt,
		Total:           rec.Total,
		OrderID:         rec.OrderID,
		UserID:          rec.UserID,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}
