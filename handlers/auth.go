package handlers

import (
	db "cyberpark/database"
	"cyberpark/database/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 獲取登入網頁
func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// 登入會員邏輯
func LoginHandler(c *gin.Context) {
	db := db.Connect()
	var user models.User
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "not founded email"})
	}
}

// 獲取註冊網頁
func SignupPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// 註冊會員邏輯
func SignupHandler(c *gin.Context) {
	db := db.Connect()
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Received user data: %+v\n", user)

	// 檢查資料庫裡是否有相同的帳號，將資料庫已存在的帳號資訊存給 existingUser 變數
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).Limit(1).Find(&existingUser).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	// 檢查是否已經存在具有相同電子郵件地址的用戶
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// 對使用者密碼進行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hash failed"})
		return
	}
	user.Password = string(hashedPassword)
	fmt.Println(user.Password)

	// 在資料庫中新增會員
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to creat user"})
		return
	}

	// 回傳成功訊息
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// func LogoutHandler(c *gin.Context) {
// 	// 處理登出請求
// }
