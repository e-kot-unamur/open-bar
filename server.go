package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var clients clientSlice

type client struct {
	conn *websocket.Conn
	send chan []byte
}

type clientSlice []client

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

	// Upgrade to a ws connection
	ws, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Keep track of the client
	client := client{ws, make(chan []byte)}
	clients = append(clients, client)
	go client.listener()
	go client.sender()

	// Send user the saved data
	all := websocketEvent{Type: "allData", All: data}
	if allByte, err := json.Marshal(all); err != nil {
		log.Println(err)
	} else {
		client.send <- allByte
	}
}

func (c client) listener() {

	// Ping pong is a heartbeat default detection system
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		// read bytes and check for error
		msgType, bytes, err := c.conn.ReadMessage() // int, []byte, error
		if err != nil {
			log.Println(err)
			// remove closed connection from clients
			for i, client := range clients {
				if c.conn == client.conn {
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
			clients.broadcast(msgType, answer)
		} else {
			// TODO ? Inform sender ?
			log.Println(err)
		}
	}
}

// All message sent will pass through this method,
// Launching this as a goroutine ensures there won't be two
// goroutine trying to write in the ws channel at the same time.
func (c client) sender() {

	ticker := time.NewTicker(pingPeriod)

	var err error
	defer log.Println(err)
	defer ticker.Stop()
	defer c.conn.Close()

	for {
		select {
		// Ping client to ensure he's still there that motherfucka
		case <-ticker.C:
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case msg := <-c.send:
			if err = c.conn.WriteMessage(1, msg); err != nil {
				return
			}
		}

	}

}

func (clients clientSlice) broadcast(msgType int, msg []byte) {
	for _, client := range clients {
		client.send <- msg
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
