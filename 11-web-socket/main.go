package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Pool struct {
	ID   string
	Conn *websocket.Conn
}

var pools []*Pool

// mengirim pesan ke semua koneksi
func broadcast(from string, pools []*Pool, msgType int, msg []byte) {
	for _, pool := range pools {
		if pool.ID != from {
			pool.Conn.WriteMessage(msgType, msg)
		}
	}
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		selfID := uuid.New().String()
		// register connection to pools
		pools = append(pools, &Pool{
			ID:   selfID,
			Conn: conn,
		})

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			reply := fmt.Sprintf("Server membalasa pesan: %s", msg)
			// Write message back to browser
			// if err = conn.WriteMessage(msgType, []byte(reply)); err != nil {
			// 	return
			// }
			broadcast(selfID, pools, msgType, []byte(reply))
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8081", nil)
}
