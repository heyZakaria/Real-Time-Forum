// const socket = new WebSocket("ws://localhost:4444/ws"); 
// ////////////////////////////////////////////////////////////////////// i will fix it from .html


// socket.onopen = function () {
//     console.log("Connected to WebSocket server");
// };

// socket.onmessage = function (event) {
//     let data = JSON.parse(event.data);
//     console.log("New Message:", data);
    
//     if (data.type === "message") {
//         let chatBox = document.getElementById("chat-box");
//         let messageElement = document.createElement("p");
//         messageElement.textContent = `${data.sender}: ${data.content}`;
//         chatBox.appendChild(messageElement);
//     }
// };

// socket.onclose = function () {
//     console.log("WebSocket connection closed");
// };

// socket.onerror = function (error) {
//     console.error("WebSocket error:", error);
// };


// //send input of user
// function sendMessage() {
//     let messageInput = document.getElementById("message-input");
//     let message = messageInput.value.trim();
    
//     if (message !== "") {
//         socket.send(JSON.stringify({ type: "message", content: message }));
//         messageInput.value = ""; 
//     }
// }





