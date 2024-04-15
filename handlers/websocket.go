package handlers

import (
	"cyberpark/database"
	"cyberpark/database/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 基礎設定
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允許所有來源的連接
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 用於定時調用 api 並發送出數據到前端
func fetchCryptoData(conn *websocket.Conn) {
	var db = database.DB

	// 1 分鐘迴圈一次
	ticker := time.NewTicker(75 * time.Second)
	defer ticker.Stop()

	for {

		// websocket 主要邏輯
		// 發送數據給前端
		var cryptoData []models.CryptoData

		if err := db.Find(&cryptoData).Error; err != nil {
			log.Println("資料庫獲取貨幣資訊失敗:", err)
			continue
		}

		err := conn.WriteJSON(cryptoData)
		if err != nil {
			log.Println("ws 傳送 JSON 失敗:", err)
			return
		}
		fmt.Println("ws 推送數據給前端")
		<-ticker.C
	}
}

func WsHomePageHandler(c *gin.Context) {
	// 將連線升級成 websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升級 ws 失敗:", err)
		return
	}

	// defer conn.Close()

	// 啟動一個 goroutine 定期發送數據到前端
	go fetchCryptoData(conn)
}
