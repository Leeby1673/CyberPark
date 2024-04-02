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

// 持有資產
type Holding struct {
	User        User    `gorm:"foreignkey:UserID"` // 定義外鍵關係
	Currency    string  // 幣種
	Amount      float64 // 金額
	Price       float64 // 幣價
	DailyChange float64 // 本日漲跌
}
