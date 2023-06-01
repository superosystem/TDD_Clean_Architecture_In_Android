package response

import "github.com/superosystem/bantumanten-backend/src/businesses/vendors"

type Vendor struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" `
	Price       uint   `json:"price"`
	VendorType  string `json:"vendor_type" `
	Address     string `json:"address"`
	Description string `json:"description" `
	Image       string `json:"image" form:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

func FromDomain(domain vendors.Domain) Vendor {
	return Vendor{
		ID:          domain.ID,
		Name:        domain.Name,
		Price:       domain.Price,
		VendorType:  domain.VendorType,
		Address:     domain.Address,
		Description: domain.Description,
		Image:       domain.Image,
		CreatedAt:   domain.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   domain.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:   domain.DeletedAt.Time.Format("02-01-2006 15:04:05"),
	}
}
