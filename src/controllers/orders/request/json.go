package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
)

type Order struct {
	Venue            string `json:"venue" validate:"required"`
	Decoration       string `json:"decoration" validate:"required"`
	Catering         string `json:"catering" validate:"required"`
	Mua              string `json:"mua" validate:"required"`
	Documentary      string `json:"documentary" validate:"required"`
	VenuePrice       uint
	DecorationPrice  uint
	CateringPrice    uint
	MuaPrice         uint
	DocumentaryPrice uint
	UserID           uint
}

func (rec *Order) ToDomain() *orders.Domain {
	return &orders.Domain{
		Venue:            rec.Venue,
		VenuePrice:       rec.VenuePrice,
		Decoration:       rec.Decoration,
		DecorationPrice:  rec.DecorationPrice,
		Catering:         rec.Catering,
		CateringPrice:    rec.CateringPrice,
		Mua:              rec.Mua,
		MuaPrice:         rec.MuaPrice,
		Documentary:      rec.Documentary,
		DocumentaryPrice: rec.DocumentaryPrice,
		UserID:           rec.UserID,
	}
}

func (rec *Order) Validate() error {
	validate := validator.New()

	err := validate.Struct(rec)

	return err
}
