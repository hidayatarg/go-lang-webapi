package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
	log.Println("Web Server Started Listening Port: 9000")
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome %s", r.URL.Path[1:])
}