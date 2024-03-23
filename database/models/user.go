package models

import (
	"time"
)

// 使用者結構
type User struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	Username  string    `gorm:"unique;size:20;column:username" form:"username"`
	Password  string    `gorm:"column:password" form:"password" binding:"required"`
	Email     string    `gorm:"unique;size:100;column:email" form:"email" binding:"required"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at"`
}
