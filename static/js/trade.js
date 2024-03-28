'use strict';

function showLimitOrderForm() {
    document.getElementById("limitOrderForm").style.display = "block";
    document.getElementById("marketOrderForm").style.display = "none";
}

function showMarketOrderForm() {
    document.getElementById("limitOrderForm").style.display = "none";
    document.getElementById("marketOrderForm").style.display = "block";
}
document.addEventListener("DOMContentLoaded", function() {


});
