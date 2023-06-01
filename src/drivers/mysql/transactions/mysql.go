package transactions

import (
	"errors"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/transactions"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) transactions.Repository {
	return &repository{
		conn: conn,
	}
}

func (r repository) Create(domain *transactions.Domain) (*transactions.Domain, error) {
	rec := FromDomain(domain)

	result := r.conn.Model(Transaction{}).Create(&rec)
	result.Last(rec)
	return rec.ToDomain(), nil
}

func (r repository) Update(ID string, domain transactions.Domain) (*transactions.Domain, error) {
	rec := FromDomain(&domain)

	err := r.conn.Model(Transaction{}).Where("id = ?", ID).Updates(&rec).Error
	if err != nil {
		return nil, err
	}
	r.conn.Model(Transaction{}).Where("id = ?", ID).First(&rec)
	return rec.ToDomain(), nil
}

func (r repository) GetAll() *[]transactions.Domain {
	var rec []Transaction

	r.conn.Model(&Transaction{}).Preload("Order").Preload("User").Find(&rec)

	var transactionsDomain []transactions.Domain

	for _, transaction := range rec {
		transactionsDomain = append(transactionsDomain, *transaction.ToDomain())
	}

	return &transactionsDomain
}

func (r repository) GetByID(ID string) (*transactions.Domain, error) {
	var rec = Transaction{}

	err := r.conn.Model(Transaction{}).Preload("Order").Preload("User").Where("id = ?", ID).Find(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}
	}

	return rec.ToDomain(), nil
}

func (r repository) GetByReferenceNumber(referenceNumber string) (*transactions.Domain, error) {
	var rec = Transaction{}

	err := r.conn.Model(Transaction{}).Preload("Order").Preload("User").Where("reference_number = ?", referenceNumber).Find(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}
	}

	return rec.ToDomain(), nil
}
