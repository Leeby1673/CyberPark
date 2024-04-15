package main

import (
	"cyberpark/database"
	"cyberpark/handlers"
	"cyberpark/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	go database.StartBackgroundTask()

	// 建立 Gin 路由器
	router := gin.Default()

	// 載入 HTML 模板
	router.LoadHTMLGlob("templates/*.html")

	// 載入 靜態文件
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	// 設定路由群組，決定中間件套用在哪些路由上
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	// 設置路由
	// 登入頁面
	router.GET("/login", handlers.LoginPageHandler)
	router.POST("/login", handlers.LoginHandler)
	// 註冊頁面
	router.GET("/signup", handlers.SignupPageHandler)
	router.POST("/signup", handlers.SignupHandler)
	// 首頁，GET 需補中間件檢查是否登入
	router.GET("/cyberpark", handlers.HomePageHandler)
	router.POST("/cyberpark", handlers.HomeHandler)
	// websocket 路由, 實現即時價格
	router.GET("/ws", handlers.WsHomePageHandler)

	// 需要中間件的路由
	// 交易買賣頁面
	authGroup.GET("/trade", handlers.TradePageHandler)
	// 交易歷史頁面
	authGroup.GET("/history", handlers.HistoryPageHandler)
	// 持有資產頁面
	authGroup.GET("/holding", handlers.HoldingPageHandler)

	// 啟動服務
	router.Run(":8080")
}
