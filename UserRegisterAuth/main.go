package main

import (
	"fmt"
	"log"
	"net/http"
	//helpers "./utils"

)

func main() {
	mux := http.NewServeMux()
	// router for login
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/", index)

	// router for register
	mux.HandleFunc("/register", register)

	log.Println("Listening Port: 8080")
	http.ListenAndServe(":8080", mux)
}

func login (w http.ResponseWriter, r *http.Request){

}

func index (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Index Page"))
}

func register (w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	pwdConfirm := r.FormValue("passwordConfirm")

	fmt.Printf("%s\n%s\n%s\n%s\n", username, email, pwd, pwdConfirm)
}