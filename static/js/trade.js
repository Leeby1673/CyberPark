'use strict';
console.log("讀取 js 成功");
// 切換現價市價顯示
function showLimitOrderForm() {
    document.getElementById("limitOrderForm").style.display = "block";
    document.getElementById("marketOrderForm").style.display = "none";
};
function showMarketOrderForm() {
    document.getElementById("limitOrderForm").style.display = "none";
    document.getElementById("marketOrderForm").style.display = "block";
};


// 總額的顯示
setupTotalCal("buyPrice", "buyQuantity", "buyTotal");
setupTotalCal("sellPrice", "sellQuantity", "sellTotal");
function setupTotalCal(priceClass, quantityClass, totalClass) {
    var pirceInput = document.querySelector("." + priceClass);
    var quantityInput = document.querySelector("." + quantityClass);

    var buyTotal = document.querySelector("." + totalClass);

    pirceInput.addEventListener("input",calculateTotal);
    quantityInput.addEventListener("input",calculateTotal);

    // 計算總額
    function calculateTotal() {
        // 獲取價格和數量的值
        var price = parseFloat(pirceInput.value);
        var quantity = parseFloat(quantityInput.value);

        // 如果價格和數量為空, 則不計算總額
        if (isNaN(price) || isNaN(quantity)) {
            buyTotal.textContent = ""; // 清空顯示
            return;
        }

        // 計算總額
        buyTotal.textContent = price * quantity;
    }
}
