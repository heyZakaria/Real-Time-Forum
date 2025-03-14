import { InitPost } from "./post.js";
import { reacPost } from "./reactPost.js";
import { comment } from "./comment.js";
import { Registred } from "./registred.js";
import { filter } from "./filter.js";
import { route } from "./router.js";
import { reactComment } from "./reactComment.js";


let submit_post = document.querySelector(".post_btn")

let post_category = document.querySelectorAll(".category")
let load_more = document.querySelector(".load_more")

export async function HandleSubmitPost() {
    let post_title = document.querySelector(".post_title")
    let post_content = document.querySelector(".post_content")
    //route(event, true)
    let userid = await Registred()
    if (!userid) {
        //// This should handled with routing
        route("/login", true)
        // window.location.replace("/login");
        return
    } else {

        let x = document.querySelector(".no_post")
        if (x != null) {
            x.remove()
        }

        let category = []

        for (let i = 0; i < post_category.length; i++) {

            if (post_category[i].checked && (post_category[i].value === "sport" || post_category[i].value === "science" || post_category[i].value === "entertainment")) {
                if (!category.includes(post_category[i].value)) {

                    category.push(post_category[i].value)
                }
            }
        }

        let z = post_title.value.trim();
        let y = post_content.value.trim();
        console.log(z);
        console.log(y);

        if (z.length == 0 || y.length == 0) {
            alert("Ensure you input a value in both fields!");

        } else {
            var OBJ = {
                id: null,
                user_id: 2,
                title: post_title.value,
                content: post_content.value,
                comment: null,
                like: 0,
                dislike: 0,
                category,
            }

            fetch("/api/addPost", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(OBJ),

            })
                .then(resp => resp.json())
                .then(data => {


                    if (data.Title != "") {

                        let posts = document.getElementById("posts")

                        data.like = 0
                        data.dislike = 0
                        let card = InitPost(data)
                        posts.insertBefore(card, posts.firstChild);

                        reacPost(data)
                        comment(data)

                    } else {
                        alert("Ensure you input a value in both fields!");

                    }
                })
        }
    }

    post_content.value = ""
    post_title.value = ""

    for (const e of post_category) {

        e.checked = false

    }

}
export async function Get_Data(url) {
    try {
        const response = await fetch(url);
        const data = await response.json();

        return data

    } catch (error) {
        console.error('Error fetching data:', error);
    }

}


let i = 0
let max = 5
let data = ""

let count = -1
/* window.onload = function () {
    Get_All_Posts()
} */

export async function Get_All_Posts() {

    let x = document.querySelector(".login_button")

    let cookie = decodeURIComponent(document.cookie.split("=")[1])
    const eqPos = cookie.indexOf('=')
    let register = document.querySelector(".register_button")

    if (cookie.length > 0) {
        // hide login ----  show logout
        let login = document.querySelector(".login_button")
        login.innerText = "Logout"
        login.style.backgroundColor = "#a00"

        if (register != null) {
            register.classList.add("hidden")
        }

        login.onclick = function () {
            if (login.innerText === "Logout") {
                // hide logout ----  show login
                // delete session
                document.cookie = 'session_id='; 'Max-Age=0'
                login.innerText = "Login"
                register.classList.remove("hidden")
                route("/login", true)
            }
        }
    }

    let postsData = await Get_Data(`/api`)

    if (postsData.posts == null) {

        let load_more_button = document.querySelector(".load_more")
        if (load_more_button) {
            load_more_button.remove()
        } else {

            let noPost = ADDElement("div", "no_post", "Be the first one!")
            noPost.style.fontWeight = "bold"
            document.querySelector(".container").appendChild(noPost)
        }
        return

    } else {
        let noPost = document.querySelector(".no_post")
        if (noPost != null) {

            noPost.remove()
        }
    }
    postsData = postsData.posts;
    data = postsData
    /* const tour = Math.floor(data.length / 5)
    
      if (tour == count) {
 
         let load_more_button = document.querySelector(".load_more")
         if (load_more_button) {
             load_more_button.remove()
         }
         return
     }
     count++ */

    if (postsData != null) {

        for (i = 0; i < postsData.length; i++) {

            // if (i < max) {

            // Initialize Post
            postsData[i].category = postsData[i].categories
            let card = InitPost(postsData[i])

            let posts = document.getElementById("posts")

            posts.appendChild(card);
            reacPost(postsData[i])
            comment(postsData[i])

            //     } else {

            //         break
            //     }
        }
        // max += 5
    }
}
// if (load_more) {

//     load_more.addEventListener("click", Get_All_Posts)
// }

export function ADDElement(elem, classs, content) {
    let posts = document.getElementById("posts")

    let x = document.createElement(elem)
    x.classList.add(classs)

    if (content != "") {
        x.innerText = content
    }

    return x
}

export function F(postsData) {

   /*  let load_more = document.querySelector(".load_more")
    load_more.style.display = 'none' */

    let posts = document.getElementById("posts")
    posts.innerHTML = ''

    for (let i = 0; i < postsData.length; i++) {

        // Initialize Post
        postsData[i].category = postsData[i].categories
        let card = InitPost(postsData[i])

        let posts = document.getElementById("posts")

        posts.appendChild(card);
        reacPost(postsData[i])

        comment(postsData[i])
        /////////// Maybe It's not working
        // reactComment(postsData[i])
    }
}

export async function updateLikeandDislike(postId, reaction, action) {

    if (reaction == "like") {

        try {
            const response = await fetch(`/api/posts/${postId}/like`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    action: action // "crimenter ola dicrementer"
                })
            });

            if (!response.ok) {
                throw new Error('Failed  e like');
            }

            return await response.json();
        } catch (error) {

            console.error('Error updating like:', error);
            throw error;
        }

    } else {

        try {
            const response = await fetch(`/api/posts/${postId}/dislike`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    action: action // "crimenter ola dicrementer"
                })
            });

            if (!response.ok) {
                throw new Error('Failed  e dislike');
            }

            return await response.json();
        } catch (error) {
            console.error('Error updating dislike:', error);
            throw error;
        }
    }
}
