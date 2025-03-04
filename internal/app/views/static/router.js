import { Get_All_Posts } from "./fetch_data.js";
import { Registred } from "./registred.js";


let home = document.querySelector(".home_button").addEventListener("click", e => {
    route(e)
})

let Login = document.querySelector(".login_button").addEventListener("click", e => {
    route(e)
})

let register = document.querySelector(".register_button").addEventListener("click", e => {
    route(e)
})

window.addEventListener('popstate', function (e) {
    // This state is comming from first param in route function ( window.history.pushState(e.target.href, null, e.target.href) )
    if (e.state) {
        console.log("Navigate Forward or Backward");
        route(e)
    }

}, false)

export function route(e, bol) {

    if (bol) {
        // This is when you try to post or react but you're not registred
        // you need to go to login
        console.log("Try To Post Or React");
        window.history.pushState("", null, e)
        handleLocation()

    } else {
        e.preventDefault()

        // we need first param to filter the case in popstate EventListener
        window.history.pushState(e.target.href, null, e.target.href)
        handleLocation()
    }
    window.scrollX = 0
    window.scrollY = 0
}


const handleLocation = async () => {

    let path = window.location.pathname
    let home = document.querySelector(".homeCode")
    let login = document.querySelector(".loginCode")
    let register = document.querySelector(".registerCode")
    let error = document.querySelector(".errorCode")

    if (path != "/register" || path != "/login" || path != "/") {
        error.classList.remove("hidden")

        register.classList.add("hidden")
        home.classList.add("hidden")
        login.classList.add("hidden")
    }

    if (path == "/") {
        home.classList.remove("hidden")

        login.classList.add("hidden")
        register.classList.add("hidden")
        error.classList.add("hidden")
        let userid = await Registred()
        if (userid) {

            Get_All_Posts();
        } else {
            route("/login", true)
        }
    }

    if (path == "/login") {

        login.classList.remove("hidden")

        home.classList.add("hidden")
        register.classList.add("hidden")
        error.classList.add("hidden")
    }

    if (path == "/register") {
        register.classList.remove("hidden")

        home.classList.add("hidden")
        login.classList.add("hidden")
        error.classList.add("hidden")
    }
}

handleLocation()
