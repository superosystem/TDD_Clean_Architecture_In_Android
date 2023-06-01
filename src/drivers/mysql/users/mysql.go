package users

import (
	"errors"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &repository{
		conn: conn,
	}
}

func (r repository) Create(domain *users.Domain) error {
	rec := FromDomain(domain)

	err := r.conn.Model(User{}).Create(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Update(ID string, domain users.Domain) (*users.Domain, error) {
	rec := FromDomain(&domain)

	err := r.conn.Model(User{}).Where("id = ?", ID).Updates(&rec).Error
	if err != nil {
		return nil, err
	}
	r.conn.Model(User{}).Where("id = ?", ID).First(&rec)
	return rec.ToDomain(), nil
}

func (r repository) Delete(ID string) bool {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetAll() *[]users.Domain {
	var rec []User

	r.conn.Model(&User{}).Find(&rec)

	var usersDomain []users.Domain

	for _, user := range rec {
		usersDomain = append(usersDomain, *user.ToDomain())
	}

	return &usersDomain
}

func (r repository) GetByID(ID string) (*users.Domain, error) {
	var rec = User{}

	err := r.conn.Model(User{}).Where("id = ?", ID).First(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrUserNotFound
		}
	}

	return rec.ToDomain(), nil
}

func (r repository) GetByEmail(email string) (*users.Domain, error) {
	var rec = User{}

	err := r.conn.Model(User{}).Where("email = ?", email).First(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrUserNotFound
		}
	}

	return rec.ToDomain(), nil
}
