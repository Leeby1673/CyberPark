package database

import (
	"cyberpark/database/models"
	"log"
)

// 獲得會員持有兩種交易貨幣的持有量、欲交易幣種的美金價
func CatchUserAsset(symbol, userEmail string) (map[string]float64, float64) {
	var user models.User
	var cryptodata models.CryptoData

	// 獲取會員可用 USDT,symbol：User -> Holding
	if err := DB.Preload("Holdings", "currency IN (?)", []string{"USDT", symbol}).
		Where("email = ?", userEmail).
		First(&user).Error; err != nil {

		log.Println("獲取會員可用 USDT 錯誤:", err)
		// 發生錯誤，返回零值和錯誤
		return nil, 0
	}
	// fmt.Println("會員的持有資產:", user.Holdings)
	// 遍歷會員持有的兩種貨幣
	userAsset := make(map[string]float64)
	for _, holding := range user.Holdings {
		userAsset[holding.Currency] = holding.Amount
	}

	// 獲取現在幣價：CryptoData
	if err := DB.Where("crypto_symbol = ?", symbol).First(&cryptodata).Error; err != nil {
		log.Println("獲取指定幣價錯誤:", err)
		// 發生錯誤，返回零值和錯誤
		return nil, 0
	}
	cryptoprice := cryptodata.CryptoPrice

	// 返回會員持有 USDT, 指定幣種的幣價
	return userAsset, cryptoprice
}
