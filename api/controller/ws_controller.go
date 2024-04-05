package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"heji-server/api/ws"
	"heji-server/domain"
	"heji-server/pkg"
	"heji-server/wsmsg"
	"log"
	"net/http"
	"sync"
)

// WSController websocket 处理
type WSController struct {
	MessageUseCase domain.MessageUseCase
}

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

func (wsc *WSController) Upgrade(ctx *gin.Context) {
	userID := ctx.GetHeader("X-User-ID")
	if userID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "User ID is required"})
		return
	}
	w := ctx.Writer
	r := ctx.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		pkg.Log.Print("upgrade:", err)
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
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var msg wsmsg.Packet
		if err := proto.Unmarshal(p, &msg); err != nil {
			fmt.Println("Error decoding Proto message:", err)
			break
		}
		ws.Receive(conn, &msg)
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
				pkg.Log.Println(err)
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
			pkg.Log.Println(err)
		}
	}
}
