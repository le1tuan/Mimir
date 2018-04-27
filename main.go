package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":1313", nil))
}
