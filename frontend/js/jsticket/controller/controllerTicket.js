import {postWithToken} from "https://jscroot.github.io/api/croot.js";
import {PostTicket,ResponseTicket} from "../config/config.js";
import {URLPostTicket} from "../template/template.js";
import {token} from '../template/template.js';

document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("formdata");
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        let data = PostTicket();
        postWithToken(URLPostTicket, 'Authorization', 'Bearer ' + token, data, ResponseTicket);
    });
});