'use strict';

document.addEventListener("DOMContentLoaded", function() {
    var statusMsg = document.getElementById("login-status")
    // 監聽註冊功能
    const loginForm = document.getElementById("loginForm");
    loginForm.addEventListener("submit", function(event) {
        event.preventDefault(); // 阻止表單的默認提交行為
        // 獲取表單數據
        const formData = new FormData(this);

        // 驗證表單的有效性並添加 CSS 類 "was-validated"
        if (!this.checkValidity()) {
            event.preventDefault();
            event.stopPropagation();
        }
        this.classList.add('was-validated');

        // 發送 POST 請求到後端服務器
        fetch("/", {
            method: "POST",
            body: formData
        })
        // 接收後端返回的響應
        .then(response => {
            if (!response.ok) {
                // 先解析 response 錯誤訊息, 才能根據不同的錯誤處理
                return response.json().then(errorData => {
                    throw new Error(errorData.error);
                });
            }
            return response.json();
        })
        .then(data => {
            // 在這裡處理後端服務器返回的響應
            console.log(data);
            // 當接收到成功的訊息，顯示 h6 標籤
            if (data.message === "Login successful"){
                    window.location.href = "/cyberpark";

            }
            // 可以根據響應中的信息執行相應的操作，例如重新導向到登錄頁面或顯示註冊成功消息
        })
        .catch(error => {
            console.log(error);
            if (error.message === "login failure") {
                statusMsg.textContent = "信箱或密碼錯誤" ;
                statusMsg.style.color = "#dc3545";
                statusMsg.style.display = 'block';
                console.error("帳號已存在, 請使用其他帳號", error);
            } else{
                console.error("在 fetch 過程發生錯誤:", error);
            }
            
            // 在這裡處理錯誤情況，例如顯示錯誤消息給用戶
        });

    });
});
