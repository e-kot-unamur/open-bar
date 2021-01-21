package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"websocket"
)

var clients []*websocket.Conn

func setupRoutes() {
	// Serve frontend build
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// Handle websocket connection
	http.HandleFunc("/ws", handleWs)
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
	listen(ws)
}

func listen(ws *websocket.Conn) {
	// each time a new
	for {
		// read bytes and check for error
		msgType, bytes, err := ws.ReadMessage() // int, []byte, error
		if err != nil {
			log.Println(err)
			continue
		}
		// convert bytes to a struct
		var message WebsocketEvent
		err = json.Unmarshal(bytes, &message)
		if err != nil {
			log.Println(err)
			continue
		}
		// handle event, and broadcast it to clients
		fmt.Println(message)
		if answer, err := handleEvent(message); err == nil {
			broadcast(clients, msgType, answer)
		} else {
			// TODO ? Inform sender ?
			log.Println(err)
		}
	}
}

func handleEvent(event WebsocketEvent) ([]byte, error) {
	var answer WebsocketEvent
	switch event.Type {
	case "newUser":
		// TODO set the ID to the number of actual users
		answer = WebsocketEvent{Type: "newUser", User: User{-1, event.Name, 0}}
	case "updateDebt":
		answer = event
	case "updatePrice":
		answer = event
	default:
		return make([]byte, 0), errors.New("Unknown type event")
	}
	return json.Marshal(answer)
}

func broadcast(clients []*websocket.Conn, msgType int, message []byte) {
	for _, client := range clients {
		if err := client.WriteMessage(msgType, message); err != nil {
			log.Println(err)
			continue
		}
	}
}
