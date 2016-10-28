package main

import (
	"net/http"
	"log"
)

func LogServer(w http.ResponseWriter, req *http.Request) {
	log.Println("", req.RequestURI)
}

func main() {
	http.HandleFunc("/", LogServer)
	log.Println("Starting server on port 80")
	http.ListenAndServe(":80", nil)
}
