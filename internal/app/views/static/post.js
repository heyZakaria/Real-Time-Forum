import { ADDElement } from "./fetch_data.js";
export function InitPost(postsData) {

   
    let card = ADDElement("div", `card_${postsData.id}`, "")
    card.style.backgroundColor = "white";
    card.style.padding = "20px";
    card.style.margin = "20px";
   
    card.style.wordWrap = "break-word";


    let card_username = ADDElement("div", `card_username`, `@${postsData.creator}`)
    card_username.style.padding = "10px"
    card_username.style.fontWeight = "bold"

    let card_title = ADDElement("div", "card_title", postsData.title)
    card_title.style.padding = "0px 10px 0px 10px"
    card_title.style.color = "#a00"
    card_title.style.fontWeight = "bold"
    card_title.style.fontSize = "23px"


    let card_content = ADDElement("div", "card_content", postsData.content)
    card_content.style.padding = "10px"
    card_content.style.fontSize = "15px"



    let like_button = ADDElement("button", `like_${postsData.id}`, `Likes ${postsData.like} `)
    like_button.style.backgroundColor = "#a00"
    like_button.style.color = "#fff"
    like_button.style.fontWeight = "bold"
    like_button.style.width = "80px"
    like_button.style.margin= "0px 10px 20px 0px"

    let dislike_button = ADDElement("button", `dislike_${postsData.id}`, `Dislikes ${postsData.dislike}`)
    dislike_button.style.backgroundColor = "#a00"
    dislike_button.style.color = "#fff"
    dislike_button.style.fontWeight = "bold"
    dislike_button.style.width = "80px"
    dislike_button.style.margin= "5px"

    card.appendChild(card_username);
    card.appendChild(card_title);
    card.appendChild(card_content);
  
    let categories
    if (postsData.category != null) {
      
        for (const cat of postsData.category) {
            
            categories = ADDElement("div", `tags`, `#${cat}`)
            categories.style.fontSize = "13px"
            categories.style.fontWeight = "bold"
            categories.style.color = "#2873ca"
            categories.style.margin= "10px"
            card.appendChild(categories)
        }

    }
  

    let comment_card = ADDElement("div", `comment_card_${postsData.id}`, "")
    
    let show_comment_button = ADDElement("button", `comment_btn_${postsData.id}`, `Show Comments`)
    show_comment_button.style.backgroundColor = "#fff"
    show_comment_button.style.color = "#a00"
    show_comment_button.style.fontWeight = "bold"
    show_comment_button.style.margin= "5px"
    show_comment_button.style.width = "110px"
    show_comment_button.style.border = "none"
    
    let comment_form = document.createElement("form");
    
    
    let comment_input = ADDElement("input", `comment_input_${postsData.id}`, "")
    comment_input.style.display = 'none'
    comment_input.style.padding = "10px"
    comment_input.style.width = "100%"

    

    // submit comment
    let submit_comment = ADDElement("button", `submit_comment_btn_${postsData.id}`, "Reply")
    submit_comment.style.display = 'none'
    submit_comment.style.margin = "10px 0px 0px 0px"
    submit_comment.style.backgroundColor = "#a00"
    submit_comment.style.color = "#fff"
    submit_comment.style.fontWeight = "bold"
    submit_comment.style.width = "80px"
    submit_comment.style.margin= "10px 10px 20px 0px"
    

    let comment_alert = ADDElement("div", `comment_alert_${postsData.id}`, "")
    comment_alert.style.display = 'none'
    comment_alert.style.color = "#a00"
    comment_alert.style.fontSize = "15px"


    card.appendChild(like_button);
    card.appendChild(dislike_button);

    comment_card.appendChild(show_comment_button);
    comment_form.appendChild(comment_input);
    comment_card.appendChild(comment_form);
    comment_card.appendChild(submit_comment);
    comment_card.appendChild(comment_alert)

    card.appendChild(comment_card)

    return card

}
