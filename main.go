package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Application starting...")
	setupRoutes()
	log.Println("Listening on :8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
