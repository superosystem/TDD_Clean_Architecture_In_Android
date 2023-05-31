package response

import "github.com/superosystem/bantumanten-backend/src/businesses/users"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Photo     string `json:"photo" form:"photo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Token struct {
	Token string `json:"token"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		FullName:  domain.FullName,
		Email:     domain.Email,
		Photo:     domain.Photo,
		CreatedAt: domain.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt: domain.DeletedAt.Time.Format("02-01-2006 15:04:05"),
	}
}
