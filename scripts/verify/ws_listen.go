package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/api/v1/ws"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("dial %s: %v", url, err)
	}
	defer conn.Close()

	if err := conn.WriteJSON(map[string]any{"type": "subscribe", "device_ids": []int64{}}); err != nil {
		log.Fatalf("subscribe: %v", err)
	}

	fmt.Println("listening for real-time pushes (Ctrl+C to stop)...")
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Println(string(data))
	}
}
