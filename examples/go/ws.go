package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func listen(conn *websocket.Conn) {
	for {
		messageType, messageContent, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("WEBSOCKET server: ", string(messageContent))

		messageResponse := []byte("WEBSOCKET server: Greetings!")
		Error := conn.WriteMessage(messageType, []byte(messageResponse))
		if err != nil {
			log.Println(Error)
			return
		}

	}
}

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("WEBSOCKET server: Client connected!")
		listen(websocket)
	})

	http.ListenAndServe(":8080", nil)
}
