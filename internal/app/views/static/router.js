import { Get_All_Posts } from "./fetch_data.js";
import { Registred } from "./registred.js";
import { myCode } from "./code.js"
import { HandleSubmitPost } from "./fetch_data.js";

import { verifyRegistration } from "./verify_registration.js";
import { verifyLogin } from "./verify_login.js";
import { filter } from "./filter.js";


let regsiter_form = document.querySelector(".regsiter_form")
let login_form = document.querySelector(".login_form")

let submit_post = document.querySelector(".post_btn")


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
        window.history.pushState({e}, null, e)
        handleLocation()

    } else {
        e.preventDefault()

        // update the param in the URL
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
    let contentWrapper = document.querySelector(".content")


    if (path != "/register" || path != "/login" || path != "/") {
        contentWrapper.innerHTML = ""
        contentWrapper.innerHTML = myCode.errata
        
        //   lunchListener(backHome)
    }

    if (path == "/") {
        contentWrapper.innerHTML = ""
        contentWrapper.innerHTML = myCode.home
        //  lunchListener(addPost, filter, loadMore)

        let userid = await Registred()
        if (userid) {

            Get_All_Posts();
            lunchListener("post_btn", "filterbutton")

        } else {
            route("/login", true)
        }
    }

    if (path == "/login") {

        contentWrapper.innerHTML = ""
        contentWrapper.innerHTML = myCode.login
        lunchListener("login_form")
    }

    if (path == "/register") {
        contentWrapper.innerHTML = ""
        contentWrapper.innerHTML = myCode.register
        lunchListener("regsiter_form")

        // Hello()
    }
}

// this handles the browser navigation buttons 
window.onpopstate = handleLocation()

//this for every URL the user try to open 
handleLocation()

function lunchListener(toListenTo, toListenTo1, toListenTo2) {
    if (toListenTo != "post_btn") {

        document.querySelector(`.${toListenTo}`).addEventListener("submit", async function (event) {
            event.preventDefault()

            if (toListenTo == "regsiter_form") {
                verifyRegistration()
                return
            }

            if (toListenTo == "login_form") {
                verifyLogin()
                return
            }

        })
    } else {
        // for submit post form
        // like dislike
        // filter
        // load more

        document.querySelector(`.${toListenTo}`).addEventListener("click", async function (event) {
            event.preventDefault()

            HandleSubmitPost()
        })
        //  reactComment()
        document.querySelector(`.${toListenTo1}`).addEventListener("click", async function (event) {
            event.preventDefault()
            filter()
        })


    }



}