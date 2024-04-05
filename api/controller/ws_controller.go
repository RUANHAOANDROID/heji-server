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
	"net/http"
	"sync"
)

// WSController websocket 处理
type WSController struct {
	MessageUseCase domain.MessageUseCase
	BillUseCase    domain.BillUseCase
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
	userID := getUserId(ctx)
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
	msgChan := make(chan wsmsg.Packet)
	go msgProcessor(msgChan, ctx, conn)
	registerHandler(wsmsg.Type_ADD_BILL, &ws.AddBillHandler{BillUseCase: wsc.BillUseCase})
	registerHandler(wsmsg.Type_DELETE_BILL, &ws.DeleteBillHandler{BillUseCase: wsc.BillUseCase})
	for {
		//解析消息 msgType 为传输类型
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			pkg.Log.Println(msgType)
			pkg.Log.Println(err)
			return
		}
		var msg wsmsg.Packet
		if err := proto.Unmarshal(p, &msg); err != nil {
			fmt.Println("Error decoding Proto message:", err)
			break
		}
		msgChan <- msg
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

// 处理器注册表
var handlers = make(map[wsmsg.Type]ws.MessageHandler)

// 注册处理器
func registerHandler(msgType wsmsg.Type, handler ws.MessageHandler) {
	handlers[msgType] = handler
}

// 处理消息的 Goroutine
func msgProcessor(incomingMessages <-chan wsmsg.Packet, ctx *gin.Context, conn *websocket.Conn) {
	for msg := range incomingMessages {
		handler, ok := handlers[msg.Type]
		if !ok {
			fmt.Println("No handler registered for message type:", msg.Type)
			continue
		}
		handler.HandleMessage(&msg, ctx, conn)
	}
}
