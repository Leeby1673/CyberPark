

// 監聽 websocket 連接的 onmessage 事件
window.socket.onmessage = function(event) {
    // 解析後端傳來的 JSON 數據
    var cryptodata = JSON.parse(event.data);
    console.log(cryptodata);

    // 在 HTML 中遍歷 JSON 數據並顯示在欄位中
    var tableBody = document.getElementById("crypto-data");
    tableBody.innerHTML = ""; // 清空表格內容

    cryptodata.forEach(function(data){
        var trRow = tableBody.insertRow(); // 在表格中插入一行, 等同 html <tr>
        var symbolCell = trRow.insertCell(0); // 等同 html <td>
        var priceCell = trRow.insertCell(1);
        var changeCell = trRow.insertCell(2);
        var marketCapCell = trRow.insertCell(3);
        var volumeCell = trRow.insertCell(4);

        // 將數據填入 <td>
        // 某些數值用 numeral.js 套件
        symbolCell.textContent = data.symbol;
        priceCell.textContent = numeral(data.price).format('$0,0.00');;
        changeCell.textContent = data.percent_change_24h.toFixed(2) + "%";
        marketCapCell.textContent = numeral(data.market_cap).format('$0,0.00');
        volumeCell.textContent = numeral(data.volume_24h).format('$0,0.00');

        // 監聽會員點擊事件，為每個幣種添加點擊事件處理器
        trRow.addEventListener("click",function(){
            // 動態生成跳轉的網址參數，根據用戶點擊的幣種
            var tradeURL = "/trade?symbol=" + data.symbol;
            // 跳轉交易頁面
            window.location.href = tradeURL;
        });
    });
};