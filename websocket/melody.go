package websocket

import (
	"fmt"

	"gopkg.in/olahol/melody.v1"
)

// Melody -- Establishing connection to a websocket.
func Melody() {
	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	fmt.Println("hello from websocket server ...")
}
