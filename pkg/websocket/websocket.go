package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)
var clientsByParticipant = make(map[string]map[*websocket.Conn]bool)
var clientsLock = &sync.Mutex{}
var Broadcast = make(chan Message)

type Message struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
	AuctionID string `json:"auction_id"`
	UserID    string `json:"user_id,omitempty"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade error:", err)
	}
	defer ws.Close()

	clientsLock.Lock()
	if clientsByParticipant[userID] == nil {
		clientsByParticipant[userID] = make(map[*websocket.Conn]bool)
	}
	clients[ws] = userID
	clientsByParticipant[userID][ws] = true
	clientsLock.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			clientsLock.Lock()
			delete(clientsByParticipant[userID], ws)
			delete(clients, ws)
			clientsLock.Unlock()
			break
		}
		Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-Broadcast
		clientsLock.Lock()
		for client := range clientsByParticipant[msg.UserID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(clientsByParticipant[msg.UserID], client)
				delete(clients, client)
			}
		}
		clientsLock.Unlock()
	}
}