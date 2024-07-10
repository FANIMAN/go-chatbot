package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ChatbotHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Chatbot endpoint called: ")
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade:", err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read:", err)
            break
        }
        log.Printf("Received: %s", msg)

        err = conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            log.Println("Write:", err)
            break
        }
    }
}
