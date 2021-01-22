package main

import (
	"log"
	"net/http"
)

// This is our app db,
// We'll save it periodically
const historyFile string = "history.json"
const port string = ":8000"

var data openBarData

func main() {
	log.Println("Starting application")

	data = load(historyFile)
	setupRoutes()
	log.Println("Up and running on port ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Server terminated.")
	}
}
