
let asocket

let sendto
let currentuser


let fetching = false
let messageloaded = 0

export function startws() {
    connectWebsocket()

    // setInterval(fetchOnlineUsers, 1000);
    setInterval(fetchOfflineUsers, 2000);

    const sendButton = document.getElementById('sendbutton')
    sendButton.addEventListener('click', sendMessage)

}

function fetchOfflineUsers() {
    fetch('/api/friends-list')
        .then(response => response.json())
        .then(users => {
            let userList = document.getElementById('friends-list');
            if (userList == null) {
                return
            }
            userList.innerHTML = ''; //clear last status

            ///////////////////////////////// OFFLINE ///////////////////////////////////// 
            for (const [key, val] of Object.entries(users.offline)) {

                if (val == currentuser) {
                    continue
                }

                let userDiv = document.createElement('div');
                userDiv.classList.add('user');

                // make random avatar 
                let img = document.createElement('img');
                //apii for random acvatr
                img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${val}`
                img.alt = val;

                let name = document.createElement('span');
                name.textContent = val;

                // if online do green 
                let status = document.createElement('div');
                status.classList.add('offline');

                userDiv.appendChild(img);
                userDiv.appendChild(name);
                userDiv.appendChild(status);
                userList.appendChild(userDiv);

            };
            let chat_title = document.getElementById("chat-title")
            let chatbox = document.getElementById("chat-box")
            let messageinput = document.getElementById("message-input")
            let sendbutton = document.getElementById("sendbutton")
            let talkingto = document.getElementById("talkingto")
            ///////////////////////////////// ONLINE /////////////////////////////////////
            for (const [key, val] of Object.entries(users.online)) {

                // console.log(user.username,"hada useer")
                if (val == currentuser) {
                    continue
                }

                let userDiv = document.createElement('div');
                userDiv.classList.add('user');

                // make random avatar 
                let img = document.createElement('img');
                //apii for random acvat
                img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${val}`
                img.alt = val;

                let name = document.createElement('span');
                name.textContent = val;


                // if online do green 
                let status = document.createElement('div');
                status.classList.add('online');
                //need to add ofline
                let showChat = {}
                showChat[val] = false
                userDiv.addEventListener('click', () => {

                    if (!showChat.val) {

                        chat_title.style.display = "block"
                        chatbox.style.display = "block"
                        messageinput.style.display = "block"
                        messageinput.addEventListener('keydown', (e) => {
                            console.log("siiiiir", e);

                        })
                        messageinput.addEventListener('keyup', () => {
                            console.log("7baassssssssss");

                        })
                        sendbutton.style.display = "block"
                        talkingto.style.display = "block"
                        sendto = val
                        talkingto.innerHTML = ""
                        talkingto.innerHTML = "Your'e talking white :" + sendto
                        initializeChat(currentuser, sendto)
                        showChat.val = true
                    } else {
                        chat_title.style.display = "none"
                        chatbox.style.display = "none"
                        messageinput.style.display = "none"
                        sendbutton.style.display = "none"
                        talkingto.style.display = "none"
                        talkingto.innerHTML = ""
                        showChat.val = false
                    }
                    //  getchatconversation()

                });

                userDiv.appendChild(img);
                userDiv.appendChild(name);
                userDiv.appendChild(status);
                userList.appendChild(userDiv);

            };
        })
        .catch(error => console.error('Error fetching ofline users:', error));
}

// setInterval(fetchoflineusers,10000);
// fetchoflineusers()

async function connectWebsocket() {

    //fetch current api from api endpoint 
    const response = await fetch("/api/current-user")
    if (!response.ok) {
        throw new Eror("not authenticated ");
    }
    //pars json response
    const data = await response.json()
    currentuser = data.username

    //user the username for websocket conection 
    //map[string]string{"username:"current user }

    asocket = new WebSocket(`ws://localhost:4444/ws?username=${data.username}`);
    try {
        asocket.onopen = function () {
            // console.log("Connected to WebSocket server");
        };

        asocket.onmessage = function (event) {
            let data = JSON.parse(event.data);

            if (data) {
                // let chatBox = document.getElementById("chat-box");

                // let messageElement = document.createElement("p");
                // messageElement.textContent = `${data.sender}: ${data.content}`;

                // chatBox.appendChild(messageElement);
                // fetchConversation()
                initializeChat(currentuser, sendto)
            }
        };

        asocket.onclose = function () {
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
            console.log("finisyo")
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
            } else {
                chatBox.appendChild(messageElement)
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

    let currenttime = new Date()

    if (message !== "") {
        let chatBox = document.getElementById("chat-box");

        let messageElement = document.createElement("p");
        messageElement.classList.add("message")
        messageElement.textContent = `${currentuser}: ${message} ${formatDate(currenttime)} "WEee`;

        chatBox.appendChild(messageElement);
        chatBox.scrollTop = chatBox.scrollHeight;

        if (sendto) {
            // let currentUser      = "dady";  // shold do seession

            // Send the message
            asocket.send(JSON.stringify({
                type: "message",
                // sender: currentUser,  //     current user session data 
                content: message,
                receiver: sendto
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

    const handleScroll = throttle(() => {
        if (chatBox.scrollTop < 30 && !fetching) {
            fetchConversation(user1, user2, true)
        }

    }, 100)
    chatBox.addEventListener("scroll", handleScroll)
}