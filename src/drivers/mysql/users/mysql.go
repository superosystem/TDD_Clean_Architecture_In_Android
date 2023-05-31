package users

import (
	"errors"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (u userRepository) Create(domain *users.Domain) error {
	rec := FromDomain(domain)

	err := u.conn.Model(User{}).Create(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Update(ID string, domain users.Domain) (*users.Domain, error) {
	rec := FromDomain(&domain)

	err := u.conn.Model(User{}).Where("id = ?", ID).Updates(&rec).Error
	if err != nil {
		return nil, err
	}
	u.conn.Model(User{}).Where("id = ?", ID).First(&rec)
	return rec.ToDomain(), nil
}

func (u userRepository) Delete(ID string) bool {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetAll() *[]users.Domain {
	var rec []User

	u.conn.Model(&User{}).Find(&rec)

	var usersDomain []users.Domain

	for _, user := range rec {
		usersDomain = append(usersDomain, *user.ToDomain())
	}

	return &usersDomain
}

func (u userRepository) GetByID(ID string) (*users.Domain, error) {
	var rec = User{}

	err := u.conn.Model(User{}).Where("id = ?", ID).First(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrUserNotFound
		}
	}

	return rec.ToDomain(), nil
}

func (u userRepository) GetByEmail(email string) (*users.Domain, error) {
	var rec = User{}

	err := u.conn.Model(User{}).Where("email = ?", email).First(&rec).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrUserNotFound
		}
	}

	return rec.ToDomain(), nil
}
