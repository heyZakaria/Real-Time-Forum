import { route } from "./router.js";

/* export function Hello() {
    */

export function verifyRegistration() {

    const errorDiv = document.getElementById('server_error');
    let username = document.getElementById("username").value;
    let email = document.getElementById("register_email").value;
    let password = document.getElementById("register_password").value;
    let confpassword = document.getElementById("register_password_2").value;

    let validusername = false
    let validemail = false
    let validpassword = false
    let validconfpassword = false

    if (username.length < 25 && username.length >= 2 && UsernameIsValid(username) == true) {
        validusername = true;
    }

    if (email.length < 60 && email.length >= 8 && EmailIsValid(email) && UsernameIsValid(email) == true) {
        validemail = true;
    }

    if (password.length < 50 && password.length > 8 && UsernameIsValid(password) == true && checkCharacter(password) == true) {
        validpassword = true;
    }

    if (password == confpassword) {
        validconfpassword = true;
    }

    if (!validusername) {
        let myerror = document.getElementById("error0");
        myerror.innerHTML = "invalid username";
        myerror.style.color = 'red';
    }

    if (!validemail) {
        let myerror = document.getElementById("error1");
        myerror.innerHTML = "invalid email";
        myerror.style.color = 'red';
    }

    if (!validpassword) {
        let myerror = document.getElementById("error2");
        myerror.innerHTML = "invalid password";
        myerror.style.color = 'red';
    }

    if (!validconfpassword) {
        let myerror = document.getElementById("error3");
        myerror.innerHTML = "invalid confpassword";
        myerror.style.color = 'red';
    }

    if (!validconfpassword || !validpassword || !validemail || !validusername) {
        return
    }

    sendHttpRequest('POST', '/register', {
        email: email,
        username: username,
        password: password,
    }).then(responseData => {

        if (!responseData.emilorusernameexsist) {
            route("/login", true)

        }
    })
        .catch(err => {
            if (err.emilorusernameexsist) {
                errorDiv.textContent = " your email or username already exists!!!"
                errorDiv.style.color = `red`

            } else if (err.InternalError) {
                errorDiv.textContent = "server problem, try later!!!"
                errorDiv.style.color = `red`

            } else {
                errorDiv.textContent = "invaliddata!!"
                errorDiv.style.color = `red`
            }
        });

}

function EmailIsValid(email) {
    const re = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return re.test(email);
}

function checkCharacter(s) {
    let hasUpper = false;
    let hasLower = false;
    let hasDigit = false;
    let hasSpecial = false;

    for (let i = 0; i < s.length; i++) {
        const ch = s[i];

        if (/[A-Z]/.test(ch)) {
            hasUpper = true;
        } else if (/[a-z]/.test(ch)) {
            hasLower = true;
        } else if (/\d/.test(ch)) {
            hasDigit = true;
        } else if (!/[a-zA-Z0-9]/.test(ch)) {
            hasSpecial = true;
        }
    }

    return hasUpper && hasLower && hasDigit && hasSpecial;
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
function UsernameIsValid(username) {
    for (let i = 0; i < username.length; i++) {
        if (username.charCodeAt(i) < 32) {
            return false;
        }
    }
    return true;
}

const sendHttpRequest = (method, url, data) => {

    const promise = new Promise((resolve, reject) => {
        const xhr = new XMLHttpRequest();
        xhr.open(method, url);

        xhr.responseType = 'json';

        if (data) {
            xhr.setRequestHeader('Content-Type', 'application/json');
        }

        xhr.onload = () => {
            if (xhr.status >= 400) {
                reject(xhr.response);

            } else {
                resolve(xhr.response);
            }
        };

        xhr.onerror = () => {
            reject('Something went wrong!');
        };

        xhr.send(JSON.stringify(data));
    });
    return promise;
};