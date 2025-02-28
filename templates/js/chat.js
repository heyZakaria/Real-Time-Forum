$(function(){

    var socket =null;
    var msgBox=$('#chatbox-text-erea')
    var message=$("#message");

    $("#chatbox").submit(function(){

        if (msg.msgBox.val())return false;
        if (!socket){
            alert("err no socket connection ")
            return false

        }

        socket.send(msgBox.val())
        msgBox.val("")
        return false


    });
   
    if (!window["WebSocket"]){

        alert("Err browser not suppor websocket")

    }else{

        socket=new WebSocket("ws://{{.Host}}/room")

        socket.onclose=function(){

            alert("conection closdd")

        }

        socket.onmessage=function(e){
            message.append($("<li>").text(e.data))
        }
    }
    





})