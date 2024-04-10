package handlers

import (
	"cyberpark/database/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

func CatchCryptoData() []models.CryptoData {

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
	// fmt.Println(string(respBody))

	// 解析 json 數據並提取需要的資訊
	var cryptoData []models.CryptoData
	gjson.ParseBytes(respBody).Get("data").ForEach(func(key, value gjson.Result) bool {
		symbol := value.Get("symbol").String()
		price := value.Get("quote.USD.price").Float()
		percentChange24H := value.Get("quote.USD.percent_change_24h").Float()
		marketCap := value.Get("quote.USD.market_cap").Float()
		volume_24h := value.Get("quote.USD.volume_24h").Float()

		cryptoData = append(cryptoData, models.CryptoData{
			Symbol:           symbol,
			Price:            price,
			PercentChange24H: percentChange24H,
			MarketCap:        marketCap,
			Volume24H:        volume_24h,
		})

		return true // 繼續迭代
	})

	for _, data := range cryptoData {
		fmt.Printf("Symbol: %s, Price: %f, PercentChange24H: %f, MarketCap: %f, Volumn24H: %f\n",
			data.Symbol, data.Price, data.PercentChange24H, data.MarketCap, data.Volume24H)
	}

	return cryptoData
}
