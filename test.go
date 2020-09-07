package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var i ironman
	var w wolverine

	mux := http.NewServeMux()

	x1 := &messageHandler{"Bu bir mesaj!"}
	x2 := &messageHandler{"Bu da yeni bir mesaj!"}

	mux.Handle("/bir", x1)
	mux.Handle("/iki", x2)

	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)
	log.Println("Listening Port: 8080")

	http.HandleFunc("/search", search)

	http.ListenAndServe(":8080", mux)
}

type ironman int

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Iron!")
}

type wolverine int

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Wolverine!")
}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(res, x.message)
}

func search(res http.ResponseWriter, r *http.Request)  {
	h1 := r.FormValue("h1")
	io.WriteString(res, "The Parameter: " + h1)
}