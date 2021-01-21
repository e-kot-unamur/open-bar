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
var data OpenBarData

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

	client, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Send user the saved data
	all := WebsocketEvent{Type: "allData", All: data}
	if message, err := json.Marshal(all); err != nil {
		log.Println(err)
		return
	} else if err := client.WriteMessage(1, message); err != nil {
		log.Println(err)
		return
	}

	// Keep track of the client
	clients = append(clients, client)
	listen(client)
}

func listen(ws *websocket.Conn) {
	// each time a new
	for {
		// read bytes and check for error
		msgType, bytes, err := ws.ReadMessage() // int, []byte, error
		if err != nil {
			log.Println(err)
			// remove closed connection from clients
			for i, client := range clients {
				if ws == client {
					clients[i] = clients[len(clients)-1]
					clients = clients[:len(clients)-1]
				}
			}
			break
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
		user := User{len(data.Users), event.Name, 0}
		answer = WebsocketEvent{Type: "newUser", User: user}
		data.Users = append(data.Users, user)
	case "updateDebt":
		answer = event
		for id, user := range data.Users {
			if user.ID == event.ID {
				data.Users[id].Debt = event.Debt
			}
		}
		fmt.Println(data.Users)
	case "updatePrice":
		answer = event
		data.Price = event.Price
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
