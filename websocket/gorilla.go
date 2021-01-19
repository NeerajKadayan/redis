package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WShandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Host", "localhost:8080")
	r.Header.Set("Connection", "Upgrade")
	r.Header.Set("Upgrade", "websocket")
	r.Header.Set("sec-WebSocket-Extensions", "extensions")
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
		fmt.Println("Hello From Gorilla!")
	}
}
