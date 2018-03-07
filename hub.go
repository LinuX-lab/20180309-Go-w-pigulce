package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type rawMsg map[string]interface{}

// Typ danych rejestracji
type msgType uint8

// Czy rejestracja/derejestracja
const (
	msgAdd msgType = iota
	msgRemove
	msgMessage
)

type message struct {
	From  string
	Data  string
	Stamp time.Time
}

type ctrlMessage struct {
	ty     msgType
	name   string
	msg    *message
	txChan *websocket.Conn
}

var db = map[string]*websocket.Conn{}
var hubChannel = make(chan ctrlMessage, 10)

func broadcast(msg rawMsg) {
	for _, client := range db {
		if err := client.WriteJSON(msg); err != nil {
			log.Println("Broadcast error", err)
		}
	}
}

func unicast(to string, msg rawMsg) {
	if client, ok := db[to]; ok {
		if err := client.WriteJSON(msg); err != nil {
			log.Println("Unicast error", err)
		}
	} else {
		log.Println("Client", to, "not found")
	}
}

func startHub() {
	go func() {
		for op := range hubChannel {

			switch op.ty {
			case msgAdd:
				// Rozgłoszenie info o nowym kliencie
				broadcast(rawMsg{"op": "add", "name": op.name})

				// Dodanie nowego klienta do bazy
				db[op.name] = op.txChan

			case msgRemove:
				// Usunięcie klienta z bazy
				delete(db, op.name)

				// Rozgłoszenie do pozostałych o zamknięciu klienta
				broadcast(rawMsg{"op": "del", "name": op.name})

			case msgMessage:
				// Rozesłanie wiadomości
				if op.msg != nil {

					if op.name == "" {
						// Albo do wszystkich ...
						broadcast(rawMsg{"op": "msg", "data": op.msg})
					} else {
						// ... albo do jednego
						unicast(op.name, rawMsg{"op": "msg", "data": op.msg})
					}
				}
			}
		}
	}()
}
