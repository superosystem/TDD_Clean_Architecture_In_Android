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

func (u userRepository) Update(ID string, domain *users.Domain) (*users.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ID string) bool {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetAll() []users.Domain {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByID(ID string) (*users.Domain, error) {
	//TODO implement me
	panic("implement me")
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

func (u userRepository) IsExistUser(ID, email string) bool {
	//TODO implement me
	panic("implement me")
}
