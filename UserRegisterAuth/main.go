package main

import (
	"fmt"
	"log"
	"net/http"
	helpers "./utils"

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

	if helpers.IsEmpty(username) ||
		helpers.IsEmpty(email) ||
		helpers.IsEmpty(pwd) ||
		helpers.IsEmpty(pwdConfirm){
		// for long in console
		log.Println("There is Empty Data")
		// result
		fmt.Fprintln(w, "There is Empty Data")
		return
	}

	// check the password
	if pwd == pwdConfirm {
		fmt.Fprintf(w, "Registration Successful!")
		// add to the db
	}else {
		fmt.Fprintf(w, "Password Does not Match")
	}
}