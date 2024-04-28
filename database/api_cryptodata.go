package database

import (
	"cyberpark/database/models"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tidwall/gjson"
)

// 固定的後台任務
func StartBackgroundTask() {
	for {
		// 5 分鐘更新一次
		go catchCryptoData()
		fmt.Println("開始更新 api 數據")
		time.Sleep(20 * time.Minute)
	}
}

// 呼叫 coinmarketcap API，將數據存到資料庫
func catchCryptoData() {
	// 建立 API 請求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 建立 url 搜尋參數
	q := url.Values{}

	// BTC = 1, ETH = 1027, USDT = 825
	q.Add("id", "1,1027,825")
	q.Add("convert", "USD")

	// 加入 Header 設定
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "835c053c-9a6a-4fec-89e9-8e2d89cd5161")
	req.URL.RawQuery = q.Encode()

	// 接收 API 回傳的 JSON
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	// defer resp.Body.Close()

	// 解析 json 數據並提取需要的資訊
	gjson.ParseBytes(respBody).Get("data").ForEach(func(key, value gjson.Result) bool {
		symbol := value.Get("symbol").String()
		price := value.Get("quote.USD.price").Float()
		percentChange24h := value.Get("quote.USD.percent_change_24h").Float()
		marketCap := value.Get("quote.USD.market_cap").Float()
		volume24h := value.Get("quote.USD.volume_24h").Float()

		// 取小數後兩位, 並四捨五入
		price = roundTwoDecimalPlaces(price)
		percentChange24h = roundTwoDecimalPlaces(percentChange24h)
		marketCap = roundTwoDecimalPlaces(marketCap)
		volume24h = roundTwoDecimalPlaces(volume24h)

		fmt.Println("小數點轉換後:", price, percentChange24h, marketCap, volume24h)

		cryptoData := models.CryptoData{
			CryptoSymbol:           symbol,
			CryptoPrice:            price,
			CryptoPercentChange24h: percentChange24h,
			CryptoMarketCap:        marketCap,
			CryptoVolume24h:        volume24h,
		}

		// 儲存到資料庫
		if err := DB.Where("crypto_symbol = ?", cryptoData.CryptoSymbol).Save(&cryptoData).Error; err != nil {
			log.Println("保存數據失敗:", err)
			return false
		}

		// 打印數據
		fmt.Printf("Symbol: %s, Price: %f, PercentChange24H: %f, MarketCap: %f, Volumn24H: %f\n",
			cryptoData.CryptoSymbol, cryptoData.CryptoPrice, cryptoData.CryptoPercentChange24h, cryptoData.CryptoMarketCap, cryptoData.CryptoVolume24h)

		return true // 繼續迭代
	})
}

// 換算函式
func roundTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
