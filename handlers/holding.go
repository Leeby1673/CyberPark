package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 獲取持有資產頁面
func HoldingPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "holding.html", nil)
}

// 實現持有資產邏輯
func HoldingHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "holding.html", nil)
	fmt.Println("實現持有資產邏輯")
}
