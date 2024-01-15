import {postWithBearer} from "https://jscroot.github.io/api/croot.js";
import {GetDataForm,ResponseUpdateTicket} from "../config/config.js";
import {token,URLChgTicket} from "../template/template.js";

document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("formdata");

    form.addEventListener("submit", function (event) {
        event.preventDefault();
        let data = GetDataForm();
        postWithBearer(URLChgTicket, token, data, ResponseUpdateTicket)
    });
});