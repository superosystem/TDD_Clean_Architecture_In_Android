package model

import "time"

type Schedule struct {
	ID     uint      `gorm:"primaryKey" json:"id"`
	Name   string    `gorm:"type:varchar(100)" json:"name"`
	Date   time.Time `json:"date"`
	Place  string    `gorm:"type:varchar(100)" json:"place"`
	UserID uint      `json:"user_id"`
	User   []User    `gorm:"many2many:schedule_user"`
}
