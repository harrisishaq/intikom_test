package entity

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"size:255;" json:"name"`
	Email     string `gorm:"size:255;unique" json:"email"`
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time `gorm:"default:null"`
}
