package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Body string
}

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Category  string    `gorm:"not null"`
}
