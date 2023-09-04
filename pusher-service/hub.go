package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    []*Client
	nextID     int
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func newHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		nextID:     0,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.onConnect(client)
		case client := <-h.unregister:
			h.onDisconnect(client)
		}
	}
}

func (h *Hub) broadcast(msg interface{}, ignore *Client) {
	data, _ := json.Marshal(msg)
	for _, c := range h.clients {
		if c != ignore {
			c.outbound <- data
		}
	}
}

func (h *Hub) send(msg interface{}, client *Client) {
	data, _ := json.Marshal(msg)
	client.outbound <- data
}

func (h *Hub) handleWebSocket(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not upgrade", http.StatusInternalServerError)
		return
	}

	client := newClient(h, socket)
	h.register <- client

	go client.write()
}

func (h *Hub) onConnect(client *Client) {
	log.Println("client connected", client.socket.RemoteAddr())

	// make new client
	h.mutex.Lock()
	defer h.mutex.Unlock()
	client.id = h.nextID
	h.nextID++
	h.clients = append(h.clients, client)
}

func (h *Hub) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	client.close()
	h.mutex.Lock()
	defer h.mutex.Unlock()

	idx := -1
	for j, c := range h.clients {
		if c.id == client.id {
			idx = j
			break
		}
	}

	// delete client
	copy(h.clients[idx:], h.clients[idx+1:])
	h.clients[len(h.clients)-1] = nil
	h.clients = h.clients[:len(h.clients)-1]

}
