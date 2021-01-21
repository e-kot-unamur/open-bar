package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

func readWebsocket(ws *websocket.Conn) {

	type Message struct {
		Event string
		Id    int
		Name  string
		Debt  int
		Price float64
	}

	for {
		msgType, message, err := ws.ReadMessage() // int, []byte, error
		if err != nil {
			log.Println(err)
			continue
		}

		var m Message
		err = json.Unmarshal(message, &m)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(m)

		// TODO save changes locally

		broadcastWebsockets(clients, msgType, message)
	}
}

func broadcastWebsockets(clients []*websocket.Conn, msgType int, message []byte) {
	for _, client := range clients {
		if err := client.WriteMessage(msgType, message); err != nil {
			log.Println(err)
			continue
		}
	}
}

func handleWs(writer http.ResponseWriter, req *http.Request) {
	log.Println("Initializing a new websocket")

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// CORS Handler, remove when quitting development
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	clients = append(clients, ws)

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
