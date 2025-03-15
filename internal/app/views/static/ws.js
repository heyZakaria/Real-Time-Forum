let asocket

let sendto
let currentuser

export function startws(){
    console.log("SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
    connectwebsocket()

    setInterval(fetchOnlineUsers, 1000);
     setInterval(fetchoflineusers,10000);

        const sendButton = document.getElementById('sendbutton')
        sendButton.addEventListener('click', sendMessage)
          


}



function fetchOnlineUsers() {

    ///need to add this to fetch_data.jsConnected



    console.log("@@@@@@@@@@@@@@@@@@@@@@@")

            fetch('/api/online-users')
                .then(response => response.json())
                .then(users => {
                    let userList = document.getElementById('online-list');
                    let chatbox=document.getElementById("chat-box")
                    let messageinput=document.getElementById("message-input")
                    let sendbutton=document.getElementById("sendbutton")
                    if (userList!=null){
                    userList.innerHTML = ''; //clear last status

                    }

 

                    
                    // users.forEach(user => {
                        for (let user  of users) {
// 
                            console.log(user.username,"hada useer")
                            if (user.username==currentuser){
                                continue
                            }
             
                        let userDiv = document.createElement('div');
                        userDiv.classList.add('user');
        
                        // make random avatar 
                        let img = document.createElement('img');
                        //apii for random acvatr
                        img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${user.username}`
                        img.alt = user.username;
        
                        let name = document.createElement('span');
                        name.textContent = user.username;

        
                        // if online do green 
                        let status = document.createElement('div');
                        status.classList.add('online');
                        //need to add ofline
                        /////////////
                    

                userDiv.addEventListener('click', () => {
                    chatbox.style.display = "block"
                    messageinput.style.display = "block"
                    sendbutton.style.display = "block"

                    console.log('selectdUser:', user.username);
                    sendto = user.username
                    fetchConversation(currentuser, sendto)
                    //  console.log("cccccccccccccccccccc",currentuser)
                    //  getchatconversation()

                });

                userDiv.appendChild(img);
                userDiv.appendChild(name);
                userDiv.appendChild(status);
                 userList.appendChild(userDiv);

 
            };
        })
        .catch(error => console.error('Error fetching online users:', error));
}
// setInterval(fetchOnlineUsers, 1000);
// fetchOnlineUsers();

        function fetchoflineusers(){
            fetch('/api/offline-users')
                .then(response => response.json())
                .then(users => {
                    let userList = document.getElementById('ofline-list');
                
                    userList.innerHTML = ''; //clear last status

                    
                        for (let user  of users) {
// 
                            console.log(user.username,"hada useer")
                            if (user.username==currentuser){
                                continue
                            }
             
                        let userDiv = document.createElement('div');
                        userDiv.classList.add('user');
        
                        // make random avatar 
                        let img = document.createElement('img');
                        //apii for random acvatr
                        img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${user.username}`
                        img.alt = user.username;
        
                        let name = document.createElement('span');
                        name.textContent = user.username;

        
                        // if online do green 
                        let status = document.createElement('div');
                        status.classList.add('ofline');

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

async function connectwebsocket() {

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
    // `4
    try {
        asocket.onopen = function () {
            console.log("Connected to WebSocket server");
        };

        asocket.onmessage = function (event) {
            let data = JSON.parse(event.data);
            console.log("New Message:", data);

            if (data) {
                // let chatBox = document.getElementById("chat-box");

                // let messageElement = document.createElement("p");
                // messageElement.textContent = `${data.sender}: ${data.content}`;


                // chatBox.appendChild(messageElement);
                // fetchConversation()
                fetchConversation(currentuser, sendto)
            }
        };

        asocket.onclose = function () {
            console.log("WebSocket connection closed");
            let chatBox = document.getElementById("chat-box");
            chatBox.innerHTML = ""; // Clear chat box

        };

    } catch (error) {

        console.log("err conct o websocket", error)
    }
}

// connectwebsocket()
async function fetchConversation(user1, user2) {
    try {
        const response = await fetch(`/api/message-history?user1=${user1}&user2=${user2}`);
        if (!response.ok) {
            throw new Error("cant fetch data");
        }

        const messages = await response.json();
        console.log("Conversation History:", messages);

        const chatBox = document.getElementById("chat-box");
        chatBox.textContent = "";

        messages.forEach(msg => {
            const messageElement = document.createElement("div");
            messageElement.classList.add("message");
            messageElement.innerHTML = `<strong>${msg.sender_id}:</strong> ${msg.message_content} <span>${formatDate(msg.created_at)}</span>`;
            chatBox.appendChild(messageElement);
            //solve scroll probleme when you add an message 
            chatBox.scrollTop = chatBox.scrollHeight;

        });

    } catch (error) {
        console.error("Error fetching conversation:", error);
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

    if (message !== "") {
        // let sendto = prompt("enter the username of the receiver:");

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
            alert("usernam!");
        }
    }
    fetchConversation(currentuser, sendto)
}
