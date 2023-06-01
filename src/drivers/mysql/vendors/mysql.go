package vendors

import (
	"errors"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) vendors.Repository {
	return &repository{
		conn: conn,
	}
}

func (r repository) Create(domain *vendors.Domain) error {
	rec := FromDomain(domain)

	err := r.conn.Model(Vendor{}).Create(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Update(ID string, domain vendors.Domain) (*vendors.Domain, error) {
	rec := FromDomain(&domain)

	err := r.conn.Model(Vendor{}).Where("id = ?", ID).Updates(&rec).Error
	if err != nil {
		return nil, err
	}
	r.conn.Model(Vendor{}).Where("id = ?", ID).First(&rec)
	return rec.ToDomain(), nil
}

func (r repository) Delete(ID string) bool {
	err := r.conn.Model(Vendor{}).Delete("id", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func (r repository) GetAll() *[]vendors.Domain {
	var rec []Vendor

	r.conn.Model(&Vendor{}).Find(&rec)

	var vendorsDomain []vendors.Domain

	for _, vendor := range rec {
		vendorsDomain = append(vendorsDomain, *vendor.ToDomain())
	}

	return &vendorsDomain
}

func (r repository) GetByID(ID string) (*vendors.Domain, error) {
	var rec = Vendor{}

	err := r.conn.Model(Vendor{}).Where("id = ?", ID).Find(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}
	}

	return rec.ToDomain(), nil
}

func (r repository) GetByName(name string) (*vendors.Domain, error) {
	var rec Vendor

	err := r.conn.Model(Vendor{}).Where("name = ?", name).Find(&rec).Error
	if err != nil {
		return nil, constant.ErrRecordNotFound
	}

	return rec.ToDomain(), nil
}

func (r repository) GetByType(vendorType string) *[]vendors.Domain {
	var rec []Vendor

	r.conn.Model(&Vendor{}).Where("vendor_type = ?", vendorType).Find(&rec)

	var vendorsDomain []vendors.Domain

	for _, vendor := range rec {
		vendorsDomain = append(vendorsDomain, *vendor.ToDomain())
	}

	return &vendorsDomain
}
