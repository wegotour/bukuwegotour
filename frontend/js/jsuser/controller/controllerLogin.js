import {postWithToken} from "https://jscroot.github.io/api/croot.js";
import {PostLogin,ResponseLogin} from "../config/config.js";
import {URLLogin} from "../template/template.js";
import {token} from '../template/template.js';

document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("formlogin");
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        let data = PostLogin();
        postWithToken(URLLogin, 'Authorization', 'Bearer ' + token, data, ResponseLogin);
    });
});