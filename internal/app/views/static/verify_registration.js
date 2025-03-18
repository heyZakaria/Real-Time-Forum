import { route } from "./router.js";

export function verifyRegistration() {

    const errorDiv = document.getElementById('server_error');
    let username = document.getElementById("username").value;

    let age = document.getElementById("age").value;
    let gender_M = document.getElementById("gender_M");
    let gender_F = document.getElementById("gender_F");
    let first_name = document.getElementById("first_name").value;
    let last_name = document.getElementById("last_name").value;
    let email = document.getElementById("register_email").value;
    let password = document.getElementById("register_password").value;
    let confpassword = document.getElementById("register_password_2").value;
    let genderIs = ""

    first_name.trim()
    last_name.trim()
    email.trim()
    password.trim()
    confpassword.trim()

    let validusername = false
    let valid_age = false
    let valid_gender_F = false
    let valid_gender_M = false
    let valid_first_name = false
    let valid_last_name = false
    let validemail = false
    let validpassword = false
    let validconfpassword = false


    if (age > 15) {
        valid_age = true
    }

    if (gender_F.checked) {
        valid_gender_F = true
        genderIs = gender_F.value
    }

    if (gender_M.checked) {
        valid_gender_M = true
        genderIs = gender_M.value
    }

    if (first_name.length >= 3 && first_name.length <= 15) {

        valid_first_name = true
    }
    if (last_name.length >= 3 && last_name.length <= 15) {
        valid_last_name = true
    }
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
        let myerror = document.getElementById("username_error");
        myerror.innerHTML = "invalid username";
        myerror.style.color = 'red';
        return
    }

    if (!valid_age) {
        let myerror = document.getElementById("age_error");
        myerror.innerHTML = "Invalid age";
        myerror.style.color = 'red';
        return
    }

    if ((!valid_gender_F && !valid_gender_M) || ((valid_gender_F && valid_gender_M))) {
        let myerror = document.getElementById("gender_error");
        myerror.innerHTML = "Invalid gender";
        myerror.style.color = 'red';
        return
    } else {
        if (valid_gender_F) {
            genderIs = "F"
        } else {
            genderIs = "M"

        }
    }

    if (!valid_first_name) {
        let myerror = document.getElementById("first_name_error");
        myerror.innerHTML = "Invalid first name";
        myerror.style.color = 'red';
        return
    }
    if (!valid_last_name) {
        let myerror = document.getElementById("last_name_error");
        myerror.innerHTML = "Invalid last name";
        myerror.style.color = 'red';
        return
    }

    if (!validemail) {
        let myerror = document.getElementById("email_error");
        myerror.innerHTML = "invalid email";
        myerror.style.color = 'red';
        return
    }

    if (!validpassword) {
        let myerror = document.getElementById("password_error");
        myerror.innerHTML = "invalid password";
        myerror.style.color = 'red';
        return
    }

    if (!validconfpassword) {
        let myerror = document.getElementById("confirme_password_error");
        myerror.innerHTML = "Invalid confirmation password";
        myerror.style.color = 'red';
        return
    }

    sendHttpRequest('POST', '/register', {
        username: username,
        age: age,
        gender: genderIs,
        first_name: first_name,
        last_name: last_name,
        email: email,
        password: password,
    }).then(responseData => {

        if (!responseData.emilorusernameexsist) {
            route("/login", true)

        }
    })
        .catch(err => {
            if (err.emilorusernameexsist) {
                errorDiv.textContent = " Your email or Username already taken!"
                errorDiv.style.color = `red`

            } else if (err.InternalError) {
                errorDiv.textContent = "Server problem, try later!"
                errorDiv.style.color = `red`

            } else {
                console.log(err, "ERRR");

                errorDiv.textContent = "Invalid data!"
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