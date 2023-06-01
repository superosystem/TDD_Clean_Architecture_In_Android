package orders

import (
	"errors"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) orders.Repository {
	return &repository{
		conn: conn,
	}
}

func (r repository) Create(domain *orders.Domain) (*orders.Domain, error) {
	rec := FromDomain(domain)

	result := r.conn.Model(Order{}).Create(&rec)
	result.Last(&rec)

	return rec.ToDomain(), nil
}

func (r repository) Update(ID string, domain orders.Domain) (*orders.Domain, error) {
	rec := FromDomain(&domain)

	err := r.conn.Model(Order{}).Where("id = ?", ID).Updates(&rec).Error
	if err != nil {
		return nil, err
	}
	r.conn.Model(Order{}).Where("id = ?", ID).First(&rec)
	return rec.ToDomain(), nil
}

func (r repository) Delete(ID string) bool {
	err := r.conn.Model(Order{}).Delete("id", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func (r repository) GetAll() *[]orders.Domain {
	var rec []Order

	r.conn.Model(&Order{}).Preload("User").Find(&rec)

	var ordersDomain []orders.Domain

	for _, order := range rec {
		ordersDomain = append(ordersDomain, *order.ToDomain())
	}

	return &ordersDomain
}

func (r repository) GetByID(ID string) (*orders.Domain, error) {
	var rec = Order{}

	err := r.conn.Model(Order{}).Preload("User").Where("id = ?", ID).Find(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}
	}

	return rec.ToDomain(), nil
}

func (r repository) GetByReferenceNumber(referenceNumber string) (*orders.Domain, error) {
	var rec = Order{}

	err := r.conn.Model(Order{}).Preload("User").Where("reference_number = ?", referenceNumber).Find(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}
	}

	return rec.ToDomain(), nil
}
