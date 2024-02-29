package entity

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"size:255;not null"`
	Description string
	Status      string `gorm:"size:50;default:'pending'"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time `gorm:"default:null"`
}
