package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

type SocketPayload struct {
	PayloadType string
	Message string
}
type SocketResponse struct {
	Message string
}
type WebSocketConnection struct {
	*websocket.Conn
	User string
}

type Rooms map[string][]*WebSocketConnection

func main(){
	liveRooms := make(Rooms, 100)
	fmt.Println(liveRooms)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		currGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			fmt.Println(err)
		}
		room := r.URL.Query().Get("room")
		if _, ok := liveRooms[room]; ok{
			liveRooms[room] = append(liveRooms[room], &WebSocketConnection{currGorillaConn, "test"})
		} else {
			liveRooms[room] = make([]*WebSocketConnection, 0)
			liveRooms[room] = append(liveRooms[room], &WebSocketConnection{currGorillaConn, "test"})			
		}
		fmt.Println(liveRooms)
	})

	http.ListenAndServe(":8080", nil)
}