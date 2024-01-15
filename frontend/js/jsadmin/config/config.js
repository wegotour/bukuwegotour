import {
    setCookieWithExpireHour
} from 'https://jscroot.github.io/cookie/croot.js';

//token API
export function getTokenFromAPI() {
    const tokenUrl = "https://asia-southeast2-wegotour-403712.cloudfunctions.net/wegotourloginadminbaru";
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

// //get dataform
// export function GetDataForm() {
//     const username = document.querySelector("#username").value;
//     const password = document.querySelector("#password").value;
//     const role = document.querySelector("#role").value;
//     console.log(password)

//     const data = {
//         username: username,
//         password: password,
//         role: role
//     };
//     return data
// }

// post login
export function PostLogin() {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const role = document.querySelector("#role").value;

    const data = {
        email: email,
        password: password,
        role: role
    };
    return data;
}

// response post login
function ResponsePostLogin(response) {
    if (response && response.token) {
        console.log('Token User:', response.token);
        setCookieWithExpireHour('Login', response.token, 2);
        window.location.href = 'https://wegotour.my.id/dashboardadmin/';
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

export function ResponseLogin(result) {
    ResponsePostLogin(result)
}