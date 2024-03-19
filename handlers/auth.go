package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登入網頁
func LoginHandler(c *gin.Context) {
	// 處理登入請求
	c.HTML(http.StatusOK, "login.html", nil)

}

// 顯示註冊網頁
func SignupPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)

}

// 註冊會員邏輯
func SignupHandler(c *gin.Context) {

}

// func LogoutHandler(c *gin.Context) {
// 	// 處理登出請求
// }
