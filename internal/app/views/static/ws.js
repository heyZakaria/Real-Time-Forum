 
let asocket

let sendto
let currentuser

 
let fetching=false
let messageloaded =0
 
export function startws() {
    connectwebsocket()

    setInterval(fetchOnlineUsers, 1000);
    setInterval(fetchOfflineUsers, 1000);

    const sendButton = document.getElementById('sendbutton')
    sendButton.addEventListener('click', sendMessage)

}

function fetchOnlineUsers() {

    ///need to add this to fetch_data.jsConnected

    fetch('/api/online-users')
        .then(response => response.json())
        .then(users => {
            let userList = document.getElementById('online-list');
            let chatbox = document.getElementById("chat-box")
            let messageinput = document.getElementById("message-input")
            let sendbutton = document.getElementById("sendbutton")
            let talkingto=document.getElementById("talkingto")

            if (userList != null) {
                userList.innerHTML = ''; //clear last status
            } else {
                return
            }

            // users.forEach(user => {
            for (let user of users) {

                // console.log(user.username,"hada useer")
                if (user.username == currentuser) {
                    continue
                }

                let userDiv = document.createElement('div');
                userDiv.classList.add('user');

                // make random avatar 
                let img = document.createElement('img');
                //apii for random acvat
                img.src = `https://api.dicebear.com/7.x/bottts/svg?seed=${user.username}`
                img.alt = user.username;

                let name = document.createElement('span');
                name.textContent = user.username;


                // if online do green 
                let status = document.createElement('div');
                status.classList.add('online');
                //need to add ofline

                userDiv.addEventListener('click', () => {
                    chatbox.style.display = "block"
                    messageinput.style.display = "block"
                    sendbutton.style.display = "block"
                    talkingto.style.display="block"

                    // console.log('selectdUser:', user.username);
                    sendto = user.username
                    talkingto.innerHTML=""
                    talkingto.innerHTML="you talk whit :"+sendto
                    initialisechat(currentuser, sendto)
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

function fetchOfflineUsers() {
    fetch('/api/offline-users')
        .then(response => response.json())
        .then(users => {
            let userList = document.getElementById('ofline-list');
            userList.innerHTML = ''; //clear last status

            for (let user of users) {

                if (user.username == currentuser) {
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
            // console.log("Connected to WebSocket server");

        };

        asocket.onmessage = function (event) {
            let data = JSON.parse(event.data);
            // console.log("New Message:", data);

            if (data) {

                // let chatBox = document.getElementById("chat-box");

                // let messageElement = document.createElement("p");
                // messageElement.textContent = `${data.sender}: ${data.content}`;


                // chatBox.appendChild(messageElement);
                // fetchConversation()
                initialisechat(currentuser, sendto)
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

 









async function fetchConversation(user1,user2,fetchmore=false) {
   console.log(messageloaded,"messageload")
    if (fetching)return
    try{

        fetching=true
        const chatBox=document.getElementById("chat-box")

        if(!fetchmore){

            chatBox.textContent=""
            messageloaded=0;
        }

    const response=await fetch(
        `/api/message-history?user1=${user1}&user2=${user2}&offset=${messageloaded}&limit=${10}`
    )
    if (!response.ok)throw new Error("cant fetch data nnnn from js")
        
    const messages=await response.json();
    if (messages==null){
        console.log("finisyo")
        return
    }

    if (messages.length===0)return

    const scrollHeight=chatBox.scrollHeight
    const scrollPosition=chatBox.scrollTop


    messages.forEach(msg => {

        const messageElement=document.createElement("div")
        messageElement.classList.add("message")
        let x=msg.sender_id+":"
        let y=msg.message_content
        let z=formatDate(msg.created_at)
        messageElement.innerText=x+ " "+ y+" "+z 

        if (fetchmore){
            chatBox.prepend(messageElement)
        }else{
            chatBox.appendChild(messageElement)
        }

    


    }


);
//messages.length=10

messageloaded+=messages.length
 
if (fetchmore) {
    chatBox.scrollTop = chatBox.scrollHeight - scrollHeight + scrollPosition;
} else {
    chatBox.scrollTop = chatBox.scrollHeight;
}




    }catch(error){
        console.log("err fetching conv",error)

    }finally{
        fetching=false
    }
    
}










/////////////////////////////////////////////////////////


















function formatDate(dateString) {
    const date = new Date(dateString);
    const options = { hour: '2-digit', minute: '2-digit', hour12: true };
    return date.toLocaleTimeString('en-US', options);
}


function sendMessage() {

    let messageInput = document.getElementById("message-input");
    let message = messageInput.value.trim();
    // let currenttime=new Date().toTimeString().split(' ')[0];   
    // let currenttime=new Date().toTimeString()
     
    
 

    if (message !== "") {
        let chatBox = document.getElementById("chat-box");

        let messageElement = document.createElement("p");
        messageElement.classList.add("message")
        messageElement.textContent = `${currentuser}: ${message} ${currenttime} `;


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
    // fetchConversation(currentuser, sendto)

                 

}


function throttle(func,delay){

    let lastcall=0
    return function(...args){
        const now=Date.now()
        if (now-lastcall<delay){
            return
        }
        lastcall=now
        return func(...args)
    }

}





function initialisechat(user1,user2){

    fetchConversation(user1,user2)

    const chatBox=document.getElementById("chat-box")

    const handleScroll=throttle(()=>{
        if (chatBox.scrollTop<30 && !fetching){
            fetchConversation(user1,user2,true)
        }

    },100)
    chatBox.addEventListener("scroll",handleScroll)



}