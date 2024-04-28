package handlers

import (
	"cyberpark/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET 獲取交易頁面
func TradePageHandler(c *gin.Context) {
	// 功能塊1, 從 Context 中獲得用戶資訊
	// 抓取上下文中的 userEmail 資訊
	user, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
	fmt.Println("會員信箱:", user)

	// 解析用戶名稱
	userEmail, ok := user.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user information"})
		return
	}
	// 將路由存進 cookie, 方便之後重新登入後可以直接導回來
	c.SetCookie("prevPath", "/trade", 1209600, "/", "localhost", false, true)

	// 功能塊2, 根據 symbol 參數，從資料料庫獲得會員可用的 USDT、幣種價格
	symbol := c.Query("symbol")

	// 進行對應的資料庫查詢
	userAmount, cryptoPrice := database.CatchUserAsset(symbol, userEmail)
	// 解值 userAmount
	usdtAmount := userAmount["USDT"]
	symbolAmount := userAmount[symbol]

	fmt.Println("USDT持有:", usdtAmount, "交易幣種持有:", symbolAmount, "交易幣種換算美金:", cryptoPrice)
	// 傳遞用戶名稱到 HTML
	c.HTML(http.StatusOK, "trade.html", gin.H{
		"userEmail":    userEmail,
		"usdtAmount":   usdtAmount,
		"symbolAmount": symbolAmount,
		"cryptoPrice":  cryptoPrice,
		"symbol":       symbol,
	})

}

// 實現交易邏輯
func TradeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "trade.html", nil)
	fmt.Println("實現交易邏輯")
}
