package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 驗證中間件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("使用中間件")
		// 檢查用戶使否登入、以及 jwt 是否成功
		if !isUserLoggedIn(c) {
			// 存儲當前頁面路由到 context
			c.Set("redirectURL", c.Request.URL.RequestURI())

			// 跳轉登入頁面
			c.Redirect(http.StatusFound, "/login")
			return
		}

		// 確定沒問題, 繼續後面的中間件或處理器
		c.Next()
	}
}

// 驗證 JWT 傳回 bool
func isUserLoggedIn(c *gin.Context) bool {
	// 從 Cookie 提取 JWT
	tokenstr, err := c.Cookie("token")
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing JWT"})
		fmt.Println("讀取 cookie 失敗")
		return false
	}

	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// 設置用於解析 jwt 的密鑰
		return []byte("your_secret_key"), nil
	})
	// 當出現錯誤，或者 token 驗證失敗，回到 login
	if err != nil || !token.Valid {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT"})
		fmt.Println("JWT 驗證失敗")
		return false
	}

	// jwt 驗證通過
	// 取出 token Subject 的數值
	userEmail, err := token.Claims.GetSubject()
	if err != nil {
		fmt.Println("讀取 sub 失敗")
		return false
	}

	fmt.Println("使用者帳號: ", userEmail)

	// jwt 驗證通過, 將用戶訊息添加到請求中
	// 可以在後續的中間件或請求處理程序中訪問 "user" 這個鍵來獲取使用者的信息
	c.Set("userEmail", userEmail)
	return true
}
