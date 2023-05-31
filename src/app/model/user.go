package model

import "time"

type User struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	FullName     string  `gorm:"type:varchar(150)" json:"full_name"`
	Email        string  `gorm:"type:varchar(100);unique;not_null" json:"email"`
	Password     string  `gorm:"type:varchar(255);not_null" json:"-"`
	PhoneNumber  string  `gorm:"type:varchar(50)" json:"phone_number"`
	ProfileImage *string `gorm:"type:longtext" json:"profile_image"`
	Role         string  `gorm:"type:enum('USER', 'ADMIN');not_null" json:"role"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
