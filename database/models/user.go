package models

import (
	"time"
)

// 會員
type User struct {
	ID           uint          // 會員辨識
	Username     string        `gorm:"unique;size:20" form:"username"`                     // 會員名稱
	Password     string        `gorm:"column:password" form:"password" binding:"required"` // 會員密碼
	Email        string        `gorm:"unique;size:100" form:"email" binding:"required"`    // 會員信箱
	Holdings     []Holding     // 會員持有資產
	Transactions []Transaction // 會員交易紀錄
	CreatedAt    time.Time     `gorm:"type:timestamp"`
	UpdatedAt    time.Time     `gorm:"type:timestamp"`
}

// 持有資產
type Holding struct {
	ID          uint
	UserID      uint      // 會員ID
	Currency    string    // 幣種
	Amount      float64   // 持有數量
	Price       float64   // 幣價換算美金
	DailyChange float64   // 本日漲跌
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}

// 交易紀錄
type Transaction struct {
	ID            uint
	UserID        uint      // 會員ID
	TradeCurrency string    // 幣種
	TradeType     string    // 交易類別（買進、賣出）
	TradeUSD      float64   // 幣價換算美金
	TradeAmount   float64   // 成交數量
	TradePrice    float64   // 成交額
	CreatedAt     time.Time `gorm:"type:timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp"`
}

// 加密貨幣
type CryptoData struct {
	Symbol           string  `gorm:"primaryKey" json:"symbol"` // 貨幣代號
	Price            float64 `json:"price"`                    // 價格
	PercentChange24h float64 `json:"percent_change_24h"`       // 24 小時漲跌
	MarketCap        float64 `json:"market_cap"`               // 市值
	Volume24h        float64 `json:"volume_24h"`               // 交易量
}
