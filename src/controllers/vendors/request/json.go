package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
)

type Vendor struct {
	Name        string `json:"name" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	VendorType  string `json:"vendor_type" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" form:"image"`
}

func (rec *Vendor) ToDomain() *vendors.Domain {
	return &vendors.Domain{
		Name:        rec.Name,
		Price:       rec.Price,
		VendorType:  rec.VendorType,
		Address:     rec.Address,
		Description: rec.Description,
		Image:       rec.Image,
	}
}

func (rec *Vendor) Validate() error {
	validate := validator.New()

	err := validate.Struct(rec)

	return err
}
