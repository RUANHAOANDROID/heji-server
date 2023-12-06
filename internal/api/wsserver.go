package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	handler2 "heji-server/internal/api/ws/handler"
	"heji-server/internal/config"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients      = make(map[*websocket.Conn]string)
	clientsMutex sync.Mutex
)

func WebSocket(router *gin.RouterGroup, config *config.Config) {
	router.GET("/ws", func(c *gin.Context) {
		handleConnections(c)
	})
}

func handleConnections(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "User ID is required"})
		return
	}
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	addConn(conn, userID)
	// 监听关闭
	conn.SetCloseHandler(func(code int, text string) error {
		conn.CloseHandler()
		conn.Close()
		err = errors.New(text)
		removeConn(conn)
		return err
	})
	//go wsWriter(ws, &writeMutex, connId)
	for {
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		handler := handler2.CreateHandler(msgType)
		handler.HandleMessage(conn, p)
	}
}
func addConn(conn *websocket.Conn, userID string) {
	clientsMutex.Lock()
	clients[conn] = userID
	clientsMutex.Unlock()
}
func removeConn(conn *websocket.Conn) {
	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()
	conn.Close()
}

func pushMessageToUser(userId string, messageType int, message []byte) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client, clientID := range clients {
		if clientID == userId {
			err := client.WriteMessage(messageType, message)
			if err != nil {
				logger.Println(err)
			}
		}
	}
}

func broadcastMessage(message map[string]interface{}) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			logger.Println(err)
		}
	}
}
