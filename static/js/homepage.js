// function getCookieValue(cookieName) {
//     var cookies = document.cookie.split(';').reduce(function (cookies, cookie) {
//         var parts = cookie.split('=');
//         cookies[parts[0].trim()] = decodeURIComponent(parts[1]);
//         return cookies;
//     }, {});
//     return cookies[cookieName];
// }

// // 獲取並顯示特定 cookie 的值
// var token = getCookieValue("token");
// if (token) {
//     document.getElementById("cookieValue").textContent = "Token: " + token;
// } else {
//     document.getElementById("cookieValue").textContent = "No username cookie found.";
// }