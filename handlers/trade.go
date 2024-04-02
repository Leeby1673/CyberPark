package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 獲取交易頁面
func TradePageHandler(c *gin.Context) {
	// 從 Context 中獲得用戶資訊
	// 懷疑！重新登入和這個有關, 應該不是了
	user, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
	// fmt.Println(user, "讀取數值")

	// 解析用戶名稱
	userEmail, ok := user.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user information"})
		return
	}
	// 將路由存進 cookie, 方便之後重新登入後可以直接導回來
	c.SetCookie("prevPath", "/trade", 1209600, "/", "localhost", false, true)

	// 傳遞用戶名稱到 HTML
	c.HTML(http.StatusOK, "trade.html", gin.H{
		"userEmail": userEmail,
	})

}

// 實現交易邏輯
func TradeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "trade.html", nil)
	fmt.Println("實現交易邏輯")
}
