



let aSocket

let sendTo
let currentUser


let fetching = false
let messageloaded = 0

export function startws() {
    connectWebsocket()

    setInterval(fetchOfflineUsers, 2000);

    const sendButton = document.getElementById('sendbutton')
    sendButton.addEventListener('click', sendMessage)

}
let messageinput
let talkedWith = []
let sortable = {}
function fetchOfflineUsers() {
    fetch('/api/friends-list')
        .then(response => response.json())
        .then(users => {
            let userList = document.getElementById('friends-list');
            if (userList == null) {
                return
            }
            userList.innerHTML = ''; //clear last status
            // let s = Object.assign(sortable, talkedWith)
            // console.log(s, "==========");
            /* for (const [key, val] of Object.entries(users)) {
                if (!sortable.keys(key)) {
                    sortable(key) = val
                }
            } */
            //console.log(users, "--------");

            ///////////////////////////////// ALL USERS ///////////////////////////////////// 
            ///////////////////////////////// KEY == USERNAME ///////////////////////////////////// 
            ///////////////////////////////// VAL == TRUE/FALSE ///////////////////////////////////// 
            for (const [key, val] of Object.entries(users)) {

                if (key == currentUser) {
                    continue
                }

                let userDiv = document.createElement('div');
                userDiv.classList.add('user');

                // make random avatar 
                let img = document.createElement('img');
                //apii for random acvatr
                img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${key}`
                img.alt = key;

                let name = document.createElement('span');
                name.textContent = key;

                // if online do green 
                let status = document.createElement('div');
                if (val) {

                    status.classList.remove('offline');
                    status.classList.add('online');
                    let showChat = {}
                    showChat[key] = false
                    userDiv.addEventListener('click', () => {

                        if (!showChat.key) {

                            chat_section.style.display = "block"
                            sendTo = key
                            talkingto.innerHTML = ""
                            talkingto.innerHTML = "Your'e talking with : " + sendTo
                            initializeChat(currentUser, sendTo)
                            showChat.key = true
                        } else {
                            chat_section.style.display = "none"

                            talkingto.innerHTML = ""
                            showChat.val = false
                            sendTo = ""
                        }
                    });

                } else {
                    /* userDiv.addEventListener('click', () => {
                        if (showChat.key) {
                            chat_section.style.display = "none"

                            talkingto.innerHTML = ""
                            showChat.val = false
                            sendTo = ""
                        }


                    }); */
                    status.classList.add('offline');
                    status.classList.remove('online');

                }


                userDiv.appendChild(img);
                userDiv.appendChild(name);
                userDiv.appendChild(status);
                userList.appendChild(userDiv);

            };
            let chat_section = document.getElementById("chat_section")
        })
        .catch(error => console.error('Error fetching ofline users:', error));
}


async function connectWebsocket() {

    //fetch current api from api endpoint 
    const response = await fetch("/api/current-user")
    if (!response.ok) {
        throw new Eror("not authenticated ");
    }
    //pars json response
    const data = await response.json()
    currentUser = data.username


    aSocket = new WebSocket(`ws://localhost:4444/ws?username=${data.username}`);
    try {
        aSocket.onopen = function () {
            console.log("Connected to WebSocket server");
            /* //if (messageinput) {
                let dots = document.querySelectorAll("#dot")
                messageinput.addEventListener('keydown', (e) => {
                    dots.forEach(dot => {
                        dot.classList.add("dot")
    
                    });
                })
                messageinput.addEventListener('keyup', () => {
                    setTimeout(() => {
                        dots.forEach(dot => {
                            dot.classList.remove("dot")
                        });
                    }, "2000");
    
                })
          //  } */
        };

        aSocket.onmessage = function (event) {
            let data = JSON.parse(event.data);

            if (data) {
                // let chatBox = document.getElementById("chat-box");

                // let messageElement = document.createElement("p");
                // messageElement.textContent = `${data.sender}: ${data.content}`;

                // chatBox.appendChild(messageElement);
                // fetchConversation()

              //  console.log("ONMESSAGE 4", data);

                let notif = document.querySelector(".notif")
                notif.innerText = "New Message"
                notif.classList.add("notification")
                setTimeout(() => {
                    notif.innerText = ""
                    notif.classList.remove("notification")
                }, "3000");
                initializeChat(currentUser, sendTo)

            }
        };

        aSocket.onclose = function () {
            // console.log("WebSocket connection closed");
            let chatBox = document.getElementById("chat-box");
            chatBox.innerHTML = ""; // Clear chat box
        };


    } catch (error) {

        console.log("err conct o websocket", error)
    }
}

async function fetchConversation(user1, user2, fetchMore = false) {
    if (fetching) return
    try {

        fetching = true
        const chatBox = document.getElementById("chat-box")

        if (!fetchMore) {

            chatBox.textContent = ""
            messageloaded = 0;
        }

        const response = await fetch(
            `/api/message-history?user1=${user1}&user2=${user2}&offset=${messageloaded}&limit=${10}`
        )
        if (!response.ok) throw new Error("cant fetch data nnnn from js")

        const messages = await response.json();
        if (messages == null) {
            return
        }

        if (messages.length === 0) return

        const scrollHeight = chatBox.scrollHeight
        const scrollPosition = chatBox.scrollTop

        messages.forEach(msg => {

            const messageElement = document.createElement("div")
            messageElement.classList.add("message")
            let x = msg.sender_id + ":"
            let y = msg.message_content
            let z = formatDate(msg.created_at)
            messageElement.innerText = x + " " + y + " " + z

            if (fetchMore) {
                chatBox.prepend(messageElement)
                //console.log("22222222222",y)
            } else {
                chatBox.prepend(messageElement)
                //console.log("1111111111111",y)

            }
        });

        messageloaded += messages.length

        if (fetchMore) {
            chatBox.scrollTop = chatBox.scrollHeight - scrollHeight + scrollPosition;
        } else {
            chatBox.scrollTop = chatBox.scrollHeight;
        }
    } catch (error) {
        console.log("err fetching conv", error)

    } finally {
        fetching = false
    }
}

function formatDate(dateString) {
    const date = new Date(dateString);
    const options = { hour: '2-digit', minute: '2-digit', hour12: true };
    return date.toLocaleTimeString('en-US', options);
}

function sendMessage() {

    let messageInput = document.getElementById("message-input");
    let message = messageInput.value.trim();

    let currentTime = new Date()

    if (message !== "") {
        let chatBox = document.getElementById("chat-box");

        let messageElement = document.createElement("p");
        messageElement.classList.add("message")
        messageElement.textContent = `${currentUser}: ${message} ${formatDate(currentTime)} `;

        chatBox.appendChild(messageElement);
        chatBox.scrollTop = chatBox.scrollHeight;

        if (sendTo) {
            /*  talkedWith.push(currentUser, sendTo)
             console.log("currentUser", currentUser);
             console.log("sendTo", sendTo);
             console.log(talkedWith); */

            console.log("SEND 1", message, sendTo);

            // Send the message
            aSocket.send(JSON.stringify({
                type: "message",
                // sender: currentUser,  //     current user session data 
                content: message,
                receiver: sendTo
            }));
            messageInput.value = "";
        } else {
            alert("Ta Select an Online user!");
        }
    }

}

function throttle(func, delay) {

    let lastcall = 0
    return function (...args) {
        const now = Date.now()
        if (now - lastcall < delay) {
            return
        }
        lastcall = now
        return func(...args)
    }

}

function initializeChat(user1, user2) {

    fetchConversation(user1, user2)

    const chatBox = document.getElementById("chat-box")
    messageinput = document.getElementById("message-input")

    const handleScroll = throttle(() => {
        if (chatBox.scrollTop < 30 && !fetching) {
            fetchConversation(user1, user2, true)
        }

    }, 100)
    chatBox.addEventListener("scroll", handleScroll)
}



async function sendnotification(sender, receiver) {



    const response = await fetch(
        `/api/sendnotification?sender=${sender}&reciver?=${receiver}`
    )

    if (!response.ok) throw new Error("cant fetx data from js sendnotification")





}