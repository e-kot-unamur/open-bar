package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Application starting...")
	setupRoutes()
	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Server terminated.")
	}
}
