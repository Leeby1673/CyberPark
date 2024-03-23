package main

import (
	"cyberpark/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 資料庫初始化
	// db := db.Connect()
	// fmt.Println("初始化資料庫", &db)

	// 建立 Gin 路由器
	router := gin.Default()

	// 載入 HTML 模板
	router.LoadHTMLGlob("templates/*.html")

	// 載入 靜態文件
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	// 設置路由
	// 登入頁面
	router.GET("/", handlers.LoginPageHandler)
	router.POST("/", handlers.LoginHandler)
	// 註冊頁面
	router.GET("/signup", handlers.SignupPageHandler)
	router.POST("/signup", handlers.SignupHandler)
	// 首頁
	router.GET("/cyberpark", handlers.HomePageHandler)

	// 啟動服務
	router.Run(":8080")
}
