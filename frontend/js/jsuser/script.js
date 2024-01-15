// functionloadrFormLogin
function showLoading() {
    $(".loader").show();
    setTimeout(function () {
        hideLoading();
    }, 10000); // 10000 milidetik = 10 detik
}

function hideLoading() {
    $(".loader").hide();
}

showLoading();

// functionloadrFormRegister
function showLoading() {
    $(".loading").show();
    setTimeout(function () {
        hideLoading();
    }, 10000); // 10000 milidetik = 10 detik
}

function hideLoading() {
    $(".loading").hide();
}

showLoading();