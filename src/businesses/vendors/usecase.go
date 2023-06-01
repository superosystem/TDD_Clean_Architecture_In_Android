package vendors

import (
	"github.com/superosystem/bantumanten-backend/src/app/constant"
)

type VendorUseCase struct {
	vendorRepository Repository
}

func NewVendorUseCase(vr Repository) UseCase {
	return &VendorUseCase{
		vendorRepository: vr,
	}
}

func (u VendorUseCase) Create(domain *Domain) error {
	var err error

	err = u.vendorRepository.Create(domain)
	if err != nil {
		return err
	}
	return nil
}

func (u VendorUseCase) Update(ID string, domain Domain) (*Domain, error) {
	_, err := u.vendorRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}

	vendorUpdated, err := u.vendorRepository.Update(ID, domain)
	if err != nil {
		return nil, err
	}

	return vendorUpdated, nil
}

func (u VendorUseCase) Delete(ID string) bool {
	return u.Delete(ID)
}

func (u VendorUseCase) GetAll() *[]Domain {
	return u.vendorRepository.GetAll()
}

func (u VendorUseCase) GetByID(ID string) (*Domain, error) {
	vendor, err := u.vendorRepository.GetByID(ID)
	if err != nil {
		if err == constant.ErrRecordNotFound {
			return nil, constant.ErrRecordNotFound
		}
		return nil, constant.ErrInternalServerError
	}
	return vendor, nil
}

func (u VendorUseCase) GetByName(name string) (*Domain, error) {
	vendor, err := u.vendorRepository.GetByName(name)
	if err != nil {
		if err == constant.ErrRecordNotFound {
			return nil, constant.ErrRecordNotFound
		}
		return nil, constant.ErrInternalServerError
	}
	return vendor, nil
}

func (u VendorUseCase) GetByType(vendorType string) *[]Domain {
	return u.vendorRepository.GetByType(vendorType)
}
