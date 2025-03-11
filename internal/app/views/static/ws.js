

let chatSocket

export function openSocket(params) {
    chatSocket = new WebSocket("ws://" + location.host + "/ws")

}
export function handleWS(msg) {

    chatSocket.send(msg)
}

export function removeElement(element) {
    document.querySelector(`.${element}`).remove()
}

function addElement(element) {

    document.querySelector(`.${element}`).remove()
}