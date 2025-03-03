import { route } from "./router";


document.querySelector(".login_form").addEventListener("submit", async function (event) {

  event.preventDefault()
  const errorDiv = document.getElementById("server_error");
  const usernameoremail = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  sendHttpRequest('POST', '/login', {
    emailorusername: usernameoremail,
    password: password,

  }).then(responseData => {

    if (responseData.isValidData) {
      route("/", true)
    }
  })
    .catch(err => {

      if (err.errserver) {
        errorDiv.innerHTML = "internal problem try later"
        errorDiv.style.color = "red"

      } else {
        errorDiv.innerHTML = "Your password or username is incorrect.!"
        errorDiv.style.color = "red"
        errorDiv.style.margin = "5px"
      }
    });

})

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
