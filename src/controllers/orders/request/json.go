package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
)

type Order struct {
	Venue       string `json:"name" validate:"required"`
	Decoration  string `json:"decoration" validate:"required"`
	Catering    string `json:"catering" validate:"required"`
	Mua         string `json:"mua" validate:"required"`
	Documentary string `json:"documentary" validate:"required"`
	UserId      uint
}

func (rec *Order) ToDomain() *orders.Domain {
	return &orders.Domain{
		Venue:       rec.Venue,
		Decoration:  rec.Decoration,
		Catering:    rec.Catering,
		Mua:         rec.Mua,
		Documentary: rec.Documentary,
		UserID:      rec.UserId,
	}
}

func (rec *Order) Validate() error {
	validate := validator.New()

	err := validate.Struct(rec)

	return err
}
