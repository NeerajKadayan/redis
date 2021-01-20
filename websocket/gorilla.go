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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func WShandler(w http.ResponseWriter, r *http.Request) {

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	reader(conn)

	fmt.Println("websocket connected ...")

}
