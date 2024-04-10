package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 獲取首頁
func HomePageHandler(c *gin.Context) {
	// 調用 CatchCryptoData 獲取幣價資訊
	cryptodata := CatchCryptoData()
	// 將幣價資訊給 html 模板
	c.HTML(http.StatusOK, "homepage.html", gin.H{"CryptoData": cryptodata})
}

// 首頁邏輯
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
}
