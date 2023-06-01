package orders

import (
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"strconv"
	"time"
)

type OrderUseCase struct {
	orderRepository  Repository
	vendorRepository vendors.Repository
	userRepository   users.Repository
}

func NewOrderUseCase(
	or Repository,
	vr vendors.Repository,
	ur users.Repository,
) UseCase {
	return &OrderUseCase{
		orderRepository:  or,
		vendorRepository: vr,
		userRepository:   ur,
	}
}

func (u OrderUseCase) Create(domain *Domain) (*Domain, error) {
	var err error
	// Create Order
	domain.ReferenceNumber = generateReferenceNumber(domain.UserID)
	domain.Status = constant.TRANSACTION_PENDING
	domain.TotalAmount = domain.VenuePrice + domain.DecorationPrice + domain.CateringPrice + domain.MuaPrice + domain.DocumentaryPrice

	order, err := u.orderRepository.Create(domain)
	if err != nil {
		return nil, err
	}

	return order, nil
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

func generateReferenceNumber(userId uint) string {
	var rf string
	var dt = time.Now().Unix()
	rf = "OR" + strconv.Itoa(int(dt)) + "BM" + strconv.Itoa(int(userId))

	return rf
}
