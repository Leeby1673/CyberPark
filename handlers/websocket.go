package handlers

import (
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
	// 1 分鐘迴圈一次
	for {
		cryptodata := CatchCryptoData()

		// websocket 主要邏輯
		// 發送數據給前端
		err := conn.WriteJSON(cryptodata)
		if err != nil {
			log.Println("ws 傳送 JSON 失敗:", err)
			return
		}

		// 計時器
		time.Sleep(70 * time.Second)
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
