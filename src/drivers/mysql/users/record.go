package users

import (
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	FullName  string         `json:"full_name" gorm:"type:varchar(100)"`
	Email     string         `json:"email" gorm:"type:varchar(100);unique" faker:"email"`
	Password  string         `json:"password" faker:"password"`
	Photo     string         `json:"photo" gorm:"type:varchar(100)"`
	Roles     string         `gorm:"type:enum('USER', 'ADMIN')" json:"roles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		FullName:  domain.FullName,
		Email:     domain.Email,
		Password:  domain.Password,
		Photo:     domain.Photo,
		Roles:     domain.Roles,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *User) ToDomain() *users.Domain {
	return &users.Domain{
		ID:        rec.ID,
		FullName:  rec.FullName,
		Email:     rec.Email,
		Password:  rec.Password,
		Photo:     rec.Photo,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
