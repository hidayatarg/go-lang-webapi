package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/userData", userData)

	http.ListenAndServe(":8080", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
	w.WriteHeader(http.StatusOK)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index Page"))
	w.WriteHeader(http.StatusOK)
}

func userData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("userInfo Page"))

	x := r.URL.Path
	w.Write([]byte(x))

}
