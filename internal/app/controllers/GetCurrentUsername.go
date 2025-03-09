package controllers

import (
	"encoding/json"
	"net/http"
)


func GetCurrentUsername(w http.ResponseWriter ,r *http.Request){

	// we already made selecuser and selectusername 

	cookie,err:=r.Cookie("session_id")

	if err!=nil{

		http.Error(w,"not authenticated",http.StatusUnauthorized)
		return
	}
	//get user id from session 
	user_id,err:=SelectUser(cookie.Value)
	if err!=nil{
		http.Error(w,"user not found",http.StatusUnauthorized)
		return 
	
	}

	//grt username 

	username,err:=SelectUsername(user_id)

	if err!=nil{
		http.Error(w,"username not fund",http.StatusInternalServerError)
		return 
	}

	// return as json

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"username":username})

    // json.NewEncoder(w),Encode(map string{{"usernameusername": username}})





}



