import { ADDElement, Get_Data } from "./fetch_data.js"
import { Registred } from "./registred.js"
import { reactComment } from "./reactComment.js";
let ShowComments = true;
let i = 0
let max = 5

export function comment(postsData) {
    // show comments and ability to add a comment

    let comment_card = document.querySelector(`.comment_card_${postsData.id}`)
    let show_comment_button = document.querySelector(`.comment_btn_${postsData.id}`)

    let card = document.querySelector(`.card_${postsData.id}`)

    let comments_wrapper = ADDElement("div", "comments_wrapper", ``)
    comments_wrapper.style.padding = "5px 0px 5px 20px";

    let comment_input = document.querySelector(`.comment_input_${postsData.id}`)
    let submit_comment = document.querySelector(`.submit_comment_btn_${postsData.id}`)
    let comment_alert = document.querySelector(`.comment_alert_${postsData.id}`)
    comment_alert.style.padding = "5px"

    // let card = document.querySelector("card")

    show_comment_button.onclick = async function RenderComment() {

        if (ShowComments) {

            comment_input.style.display = 'block'
            comment_input.placeholder = `Jawbo ma tskootch lih`;
            comment_input.style.width = "400px"
            submit_comment.style.display = 'block'
            show_comment_button.innerText = "Hide Comments"
            ShowComments = false


            submit_comment.onclick = async function Comment() {
                let userid = await Registred()
                if (!userid) {
                    window.location.replace("/login");
                } else {
                    let x = comment_input.value.trim();

                    if (x.length != 0) {

                        var OBJ = {
                            content: comment_input.value,
                            post_id: `${postsData.id}`,
                            user_id: 5,
                        }

                        fetch("/api/addComment", {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(OBJ),

                        })
                            .then(resp => resp.json())
                            .then(data => {

                                comment_alert.style.display = 'none'
                                // ADD COMMENTED BY ????????????
                                let comment = ADDElement("div", "comment", comment_input.value)
                                comment.style.width = "400px"


                                let comment_like_button = ADDElement("button", `comment_like_${data.id}`, `Likes ${data.like} `)


                                comment_like_button.style.backgroundColor = "#a00"
                                comment_like_button.style.color = "#fff"
                                comment_like_button.style.fontWeight = "bold"
                                comment_like_button.style.width = "80px"
                                comment_like_button.style.margin = "0px 10px 20px 0px"

                                let comment_dislike_button = ADDElement("button", `comment_dislike_${data.id}`, `Dislikes ${data.dislike}`)
                                comment_dislike_button.style.backgroundColor = "#a00"
                                comment_dislike_button.style.color = "#fff"
                                comment_dislike_button.style.fontWeight = "bold"
                                comment_dislike_button.style.width = "80px"
                                comment_dislike_button.style.margin = "5px"

                                let comment_react_buttons_wrapper = ADDElement("div", `comment_react_buttons_wrapper`, ``)

                                comment_react_buttons_wrapper.appendChild(comment_like_button)
                                comment_react_buttons_wrapper.appendChild(comment_dislike_button)
                                comment.appendChild(comment_react_buttons_wrapper)






                                //comments_wrapper.appendChild(comment);
                                comments_wrapper.insertBefore(comment, comments_wrapper.firstChild);
                                comment_input.value = ""

                                reactComment(data)

                            })

                    } else {
                        comment_alert.style.display = 'block'

                        comment_alert.innerText = "Empty input"

                        comment_card.insertBefore(comment_alert, comment_card.firstChild);

                    }

                }

            }

            let ALLData = await Get_Data(`/api`)

            let data = ALLData.posts;

            for (let i = 0; i < data.length; i++) {


                if (data[i].id == postsData.id) {
                    if (data[i].comment != null) {

                        for (let j = 0; j < data[i].comment.length; j++) {
                            // ADD COMMENTED BY ????????????
                            ///////////////////////////
                            let Comm = ADDElement("div", "comment", `${data[i].comment[j].content}`)
                            Comm.style.padding = "3px"
                            Comm.style.width = "400px"
                            Comm.style.display = "flex"
                            Comm.style.flexDirection = "column"


                            let comment_like_button = ADDElement("button", `comment_like_${data[i].comment[j].id}`, `Likes ${data[i].comment[j].like} `)


                            comment_like_button.style.backgroundColor = "#a00"
                            comment_like_button.style.color = "#fff"
                            comment_like_button.style.fontWeight = "bold"
                            comment_like_button.style.width = "80px"
                            comment_like_button.style.margin = "0px 10px 20px 0px"

                            let comment_dislike_button = ADDElement("button", `comment_dislike_${data[i].comment[j].id}`, `Dislikes ${data[i].comment[j].dislike}`)
                            comment_dislike_button.style.backgroundColor = "#a00"
                            comment_dislike_button.style.color = "#fff"
                            comment_dislike_button.style.fontWeight = "bold"
                            comment_dislike_button.style.width = "80px"
                            comment_dislike_button.style.margin = "5px"

                            let comment_react_buttons_wrapper = ADDElement("div", `comment_react_buttons_wrapper`, ``)

                            comment_react_buttons_wrapper.appendChild(comment_like_button)
                            comment_react_buttons_wrapper.appendChild(comment_dislike_button)
                            Comm.appendChild(comment_react_buttons_wrapper)


                            comments_wrapper.appendChild(Comm);

                        }

                    } else {

                    }

                    break
                }
            }
            comment_card.appendChild(comments_wrapper)


            for (let i = 0; i < data.length; i++) {
                if (Array.isArray(data[i].comment)) {

                    for (let j = 0; j < data[i].comment.length; j++) {
                        reactComment(data[i].comment[j])
                    }
                }
            }

        } else {

            while (comments_wrapper.firstChild) {
                comments_wrapper.firstChild.remove()

            }
            comment_alert.style.display = 'none'

            comment_input.value = ""

            submit_comment.style.display = 'none'
            comment_input.style.display = 'none'
            show_comment_button.innerText = "Show Comments"

            ShowComments = true


        }

        card.appendChild(show_comment_button)


    }
}