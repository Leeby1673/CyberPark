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
	User        User      `gorm:"foreignkey:UserID"` // 定義外鍵關係
	Currency    string    // 幣種
	Amount      float64   // 金額
	Price       float64   // 幣價
	DailyChange float64   // 本日漲跌
	CreatedAt   time.Time `gorm:"type:timestamp;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;column:updated_at"`
}

// 加密貨幣
type CryptoData struct {
	Symbol           string  `json:"symbol"`             // 貨幣代號
	Price            float64 `json:"price"`              // 價格
	PercentChange24H float64 `json:"percent_change_24h"` // 24 小時漲跌
	MarketCap        float64 `json:"market_cap"`         // 市值
	Volume24H        float64 `json:"volume_24H"`         // 交易量
}
