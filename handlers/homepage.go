package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 獲取首頁
func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
}

// 首頁邏輯
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
}
