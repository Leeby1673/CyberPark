package handlers

import (
	db "cyberpark/database"
	"cyberpark/database/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 獲取登入網頁 處理器
func LoginPageHandler(c *gin.Context) {
	// 獲取
	referer := c.Request.Header.Get("referer")
	c.HTML(http.StatusOK, "login.html", gin.H{"referer": referer})
}

// 登入會員邏輯 處理器
func LoginHandler(c *gin.Context) {
	db := db.Connect()
	var user models.User
	var existingUser models.User

	// 使用者輸入的資料映射到 user
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查資料庫有無使用者的帳號
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "login failure"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login failure"})
		return
	}

	// 產生 JWT
	token, err := generateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	// 將 JWT 放在 cookie, cookie 時間是 13 天, 略短 JWT
	c.SetCookie("token", token, 1123200, "/", "localhost", false, true)

	// 讀取 cookie, 找回原路使用
	prevPath, err := c.Cookie("prevPath")
	fmt.Println(err)

	// 若出錯, 表示首次登入
	if err != nil {
		prevPath = "/cyberpark"
	}

	// 刪除路由 cookie, 傳送資訊給前端
	c.SetCookie("prevPath", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"redirect": prevPath,
		"token":    token,
	})
}

// 獲取註冊網頁 處理器
func SignupPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// 註冊會員邏輯 處理器
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

// 生產 JWT func
func generateJWT(email string) (string, error) {
	// 定義簽署 JWT 的密鑰
	jwtkey := []byte("your_secret_key")

	// 使用聲明和密鑰建立 token, JWT 時間是 1 分鐘
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// sub 代表用戶, 時間設定為 2個禮拜
		"exp": jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		"sub": email,
	})
	tokenstring, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
