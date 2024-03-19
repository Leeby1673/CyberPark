package main

import (
	db "cyberpark/database"
	"cyberpark/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 資料庫初始化
	db := db.Connect()
	fmt.Println("初始化資料庫", &db)

	// 建立 Gin 路由器
	router := gin.Default()

	// 載入 HTML 模板
	router.LoadHTMLGlob("templates/*.html")

	// 載入 靜態文件
	router.Static("/static", "./static")

	// 設置路由
	router.GET("/", handlers.LoginHandler)
	router.GET("/signup", handlers.SignupPageHandler)
	router.POST("/signup", handlers.SignupHandler)

	// 啟動服務
	router.Run(":8080")
}
