package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ws", handleWs)
}

func handleWs(writer http.ResponseWriter, req *http.Request) {
	log.Println("Initializing a new websocket")

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// FIXME : CORS Handler, remove when quitting development
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	client, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Send user the saved data
	all := websocketEvent{Type: "allData", All: data}
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
		var message websocketEvent
		err = json.Unmarshal(bytes, &message)
		if err != nil {
			log.Println(err)
			continue
		}
		// handle event, and broadcast it to clients
		if answer, err := handleEvent(message); err == nil {
			broadcast(clients, msgType, answer)
		} else {
			// TODO ? Inform sender ?
			log.Println(err)
		}
	}
}

func handleEvent(event websocketEvent) ([]byte, error) {
	log.Println("New event received ", event.Type)
	var answer websocketEvent
	switch event.Type {
	case "newUser":
		user := userData{len(data.Users), event.Name, 0}
		answer = websocketEvent{Type: "newUser", User: user}
		data.Users = append(data.Users, user)
	case "updateDebt":
		answer = event
		for id, user := range data.Users {
			if user.ID == event.ID {
				data.Users[id].Debt = event.Debt
			}
		}
	case "updatePrice":
		answer = event
		data.Price = event.Price
	default:
		return make([]byte, 0), errors.New("Unknown type event")
	}
	save(data, historyFile)
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
