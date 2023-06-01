package response

import "github.com/superosystem/bantumanten-backend/src/businesses/orders"

type Order struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	ReferenceNumber  string `json:"reference_number"`
	TotalAmount      uint   `json:"total_amount"`
	Venue            string `json:"venue" `
	VenuePrice       uint   `json:"venue_price" `
	Decoration       string `json:"decoration" `
	DecorationPrice  uint   `json:"decoration_price" `
	Catering         string `json:"catering" `
	CateringPrice    uint   `json:"catering_price"`
	Mua              string `json:"mua"`
	MuaPrice         uint   `json:"mua_price"`
	Documentary      string `json:"documentary"`
	DocumentaryPrice uint   `json:"documentary_price"`
	UserID           uint   `json:"user_id"`
}

func FromDomain(domain orders.Domain) Order {
	return Order{
		ID:               domain.ID,
		ReferenceNumber:  domain.ReferenceNumber,
		TotalAmount:      domain.TotalAmount,
		Venue:            domain.Venue,
		VenuePrice:       domain.VenuePrice,
		Decoration:       domain.Decoration,
		DecorationPrice:  domain.DecorationPrice,
		Catering:         domain.Catering,
		CateringPrice:    domain.CateringPrice,
		Mua:              domain.Mua,
		MuaPrice:         domain.MuaPrice,
		Documentary:      domain.Documentary,
		DocumentaryPrice: domain.DocumentaryPrice,
		UserID:           domain.UserID,
	}
}
