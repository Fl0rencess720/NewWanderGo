package Util

import (
	"log"

	"github.com/gorilla/websocket"
)

var wsPool = make(map[string]*websocket.Conn)

func AddConnection(userAccount string, conn *websocket.Conn) {
	wsPool[userAccount] = conn
}

func GetConnection(userAccount string) (*websocket.Conn, bool) {
	conn, ok := wsPool[userAccount]
	return conn, ok
}

func RemoveConnection(userAccount string) {
	delete(wsPool, userAccount)
}

func BroadcastMessage(userAccount string, message []byte) {
	conn, ok := GetConnection(userAccount)
	if ok {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
		}
	}
}
