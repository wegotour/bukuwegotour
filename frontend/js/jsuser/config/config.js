import {
    setCookieWithExpireHour
} from 'https://jscroot.github.io/cookie/croot.js';

//token api
export function getTokenFromAPI() {
    const tokenUrl = "https://asia-southeast2-wegotour-403712.cloudfunctions.net/wegotourlogin";
    fetch(tokenUrl)
        .then(response => response.json())
        .then(tokenData => {
            if (tokenData.token) {
                userToken = tokenData.token;
                console.log('Token dari API:', userToken);
            }
        })
        .catch(error => console.error('Gagal mengambil token:', error));
}

//get data 
export function GetDataForm() {
    const username = document.querySelector("#username").value;
    const password = document.querySelector("#password").value;
    const role = document.querySelector("#role").value;
    console.log(password)

    const data = {
        username: username,
        password: password,
        role: role
    };
    return data
}

// post login
export function PostLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const role = document.getElementById("role").value;

    const data = {
        username: username,
        password: password,
        role: role
    };
    return data;
}

// alert post 
export function AlertPost(value) {
    alert(value.message + "\nRegistrasi Berhasil")
    window.location.href = "index.html"
}

// alert update
export function AlertUpdate(value) {
    alert(value.message + "\nUpdate Berhasil")
    window.location.href = "index.html"
}

// response post login
function ResponsePostLogin(response) {
    if (response && response.token) {
        console.log('Token User:', response.token);
        setCookieWithExpireHour('Login', response.token, 2);
        window.location.href = 'https://wegotour.my.id/dashboard/';
        Swal.fire({
            icon: 'success',
            title: 'Login Successful',
            text: 'You have successfully logged in!',
          });
} else {
  Swal.fire({
            icon: 'error',
            title: 'Login Failed',
            text: 'Invalid email or password. Please try again.',
          });
}
}


export function ResponsePost(result) {
    AlertPost(result);
}

export function ResponseUpdate(result) {
    AlertUpdate(result);
}

export function ResponseLogin(result) {
    ResponsePostLogin(result)
}