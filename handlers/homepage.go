package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
}
