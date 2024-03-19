package models

import (
	"time"
)

// 使用者結構
type User struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	Username  string    `gorm:"unique;size:20;column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"unique;size:100;column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
