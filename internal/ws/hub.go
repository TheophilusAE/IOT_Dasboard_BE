package ws

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	conn   *websocket.Conn
	send   chan []byte
	filter map[int64]bool
}

type Hub struct {
	mu      sync.RWMutex
	clients map[*client]struct{}
}

func NewHub() *Hub {
	return &Hub{clients: make(map[*client]struct{})}
}

func (h *Hub) register(c *client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c] = struct{}{}
}

func (h *Hub) unregister(c *client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[c]; ok {
		delete(h.clients, c)
		close(c.send)
	}
}

func (h *Hub) Broadcast(msg OutgoingMessage, deviceID *int64) {
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Printf("ws: failed to marshal broadcast message: %v", err)
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()
	for c := range h.clients {
		if deviceID != nil && c.filter != nil && !c.filter[*deviceID] {
			continue
		}
		select {
		case c.send <- payload:
		default:
			log.Printf("ws: client send buffer full, dropping message")
		}
	}
}

const (
	writeWait  = 10 * time.Second
	pingPeriod = 30 * time.Second
	sendBuffer = 32
)

func (h *Hub) ServeClient(conn *websocket.Conn) {
	c := &client{conn: conn, send: make(chan []byte, sendBuffer)}
	h.register(c)

	go h.writePump(c)
	h.readPump(c)
}

func (h *Hub) readPump(c *client) {
	defer func() {
		h.unregister(c)
		c.conn.Close()
	}()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		var sub SubscribeMessage
		if err := json.Unmarshal(data, &sub); err != nil {
			continue
		}
		if sub.Type != "subscribe" {
			continue
		}

		if len(sub.DeviceIDs) == 0 {
			c.filter = nil
		} else {
			c.filter = make(map[int64]bool, len(sub.DeviceIDs))
			for _, id := range sub.DeviceIDs {
				c.filter[id] = true
			}
		}

		ack, _ := json.Marshal(OutgoingMessage{Type: TypeSubscribed, Data: sub.DeviceIDs})
		select {
		case c.send <- ack:
		default:
		}
	}
}

func (h *Hub) writePump(c *client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case data, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		case <-ticker.C:

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
