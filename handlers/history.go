package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 獲取交易歷史頁面
func HistoryPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "history.html", nil)
}

// 實現交易歷史邏輯
func HistoryHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "history.html", nil)
	fmt.Println("實現交易歷史邏輯")
}
