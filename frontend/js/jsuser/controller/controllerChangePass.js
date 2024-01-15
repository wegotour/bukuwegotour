import {postWithBearer} from "https://jscroot.github.io/api/croot.js";
import {GetDataForm,ResponseUpdate} from "../config/config.js";
import {token,URLChgPass} from "../template/template.js";

document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");

    form.addEventListener("submit", function (event) {
        event.preventDefault();
        let data = GetDataForm();
        postWithBearer(URLChgPass, token, data, ResponseUpdate)
    });
});