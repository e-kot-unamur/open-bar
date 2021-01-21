package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func readWebsocket(ws *websocket.Conn) {
	for {
		msgType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// TODO remove me
		fmt.Println(string(message))

		if err := ws.WriteMessage(msgType, message); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleWs(writer http.ResponseWriter, req *http.Request) {
	log.Println("Initializing a new websocket")

	// CORS Handler, remove when quitting development
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println(err)
	}

	readWebsocket(ws)
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ws", handleWs)
}

func main() {
	log.Println("Application starting...")
	setupRoutes()
	log.Println("Listening on :8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
