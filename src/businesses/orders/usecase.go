package orders

import (
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"strconv"
	"time"
)

type OrderUseCase struct {
	orderRepository  Repository
	vendorRepository vendors.Repository
}

func NewOrderUseCase(
	or Repository,
	vr vendors.Repository,
) UseCase {
	return &OrderUseCase{
		orderRepository:  or,
		vendorRepository: vr,
	}
}

func (u OrderUseCase) Create(domain *Domain) error {
	var err error
	// Generate RF
	domain.ReferenceNumber = generateReferenceNumber(domain.UserID)
	// Setup Price
	domain.VenuePrice, err = u.getVendorPrice(domain.Venue)
	domain.DecorationPrice, err = u.getVendorPrice(domain.Decoration)
	domain.CateringPrice, err = u.getVendorPrice(domain.Catering)
	domain.MuaPrice, err = u.getVendorPrice(domain.Mua)
	domain.DocumentaryPrice, err = u.getVendorPrice(domain.Documentary)
	if err != nil {
		return err
	}
	domain.TotalAmount = domain.VenuePrice + domain.DecorationPrice + domain.CateringPrice + domain.MuaPrice + domain.DocumentaryPrice
	// Save Order
	err = u.orderRepository.Create(domain)
	if err != nil {
		return err
	}

	return nil
}

func (u OrderUseCase) Update(ID string, domain Domain) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (u OrderUseCase) Delete(ID string) bool {
	//TODO implement me
	panic("implement me")
}

func (u OrderUseCase) GetAll() *[]Domain {
	return u.orderRepository.GetAll()
}

func (u OrderUseCase) GetByID(ID string) (*Domain, error) {
	order, err := u.orderRepository.GetByID(ID)
	if err != nil {
		if err == constant.ErrRecordNotFound {
			return nil, constant.ErrRecordNotFound
		}
		return nil, constant.ErrInternalServerError
	}
	return order, nil
}

func (u OrderUseCase) GetByReferenceNumber(referenceNumber string) (*Domain, error) {
	order, err := u.orderRepository.GetByReferenceNumber(referenceNumber)
	if err != nil {
		if err == constant.ErrRecordNotFound {
			return nil, constant.ErrRecordNotFound
		}
		return nil, constant.ErrInternalServerError
	}
	return order, nil
}

func (u OrderUseCase) getVendorPrice(name string) (uint, error) {
	vendor, err := u.vendorRepository.GetByName(name)
	if err != nil {
		return 0, err
	}
	return vendor.Price, nil
}

func generateReferenceNumber(userId uint) string {
	var rf string
	var dt = time.Now().Unix()
	rf = "OR" + strconv.Itoa(int(dt)) + "BM" + strconv.Itoa(int(userId))

	return rf
}
