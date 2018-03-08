package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Identyfikatorem jest IP klienta połączenia
	ID := r.RemoteAddr

	// Zarejestruj go w systemie
	hubChannel <- &ctrlMessage{
		ty:     msgAdd,
		name:   ID,
		txChan: conn,
	}

	// Jak funkcja się skończy (pętla poniżej), to rozrejestruj klienta
	defer func() {
		hubChannel <- &ctrlMessage{
			ty:   msgRemove,
			name: ID,
		}
	}()

	for {
		// Pętla pobierająca komunikaty od klienta
		var msg rawMsg
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		if op, found := msg["op"]; found {
			switch op {
			case "rename":
			case "msg":
			}
		}

	}
}
