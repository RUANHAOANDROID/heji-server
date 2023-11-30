package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type WebSocketHub struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

var hub = WebSocketHub{
	clients: make(map[*websocket.Conn]bool),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	hub.mutex.Lock()
	hub.clients[conn] = true
	hub.mutex.Unlock()

	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			hub.mutex.Lock()
			delete(hub.clients, conn)
			hub.mutex.Unlock()
			break
		}

		// 在实际应用中，根据接收到的消息进行相应的处理，例如更新账本、广播给其他用户等
		fmt.Printf("Received message from client: %+v\n", message)
	}
}

func broadcastMessage(message map[string]interface{}) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	for client := range hub.clients {
		err := client.WriteJSON(message)
		if err != nil {
			fmt.Println(err)
			delete(hub.clients, client)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go http.ListenAndServe(":8080", nil)

	// 示例：模拟一个记账操作，向所有订阅 "accounting" 主题的用户广播
	go func() {
		for i := 1; i <= 5; i++ {
			message := map[string]interface{}{
				"topic":     "accounting",
				"operation": fmt.Sprintf("Transaction %d", i),
			}
			broadcastMessage(message)
		}
	}()

	select {}
}
