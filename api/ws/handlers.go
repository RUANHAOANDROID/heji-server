package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"heji-server/wsmsg"
)

// 定义全局处理器变量
var handler AckHandler

func init() {
	handler = AckHandler{}
	addBillHandler := &AddBillHandler{}
	handler.SetNext(addBillHandler)
}
func Receive(conn *websocket.Conn, msg *wsmsg.Packet) {
	handler.HandleMessage(conn, msg)
}

type MessageHandler interface {
	HandleMessage(conn *websocket.Conn, msg *wsmsg.Packet)
}
type AckHandler struct {
	Next MessageHandler
}

func (h *AckHandler) SetNext(handler MessageHandler) {
	h.Next = handler
}

func (h *AckHandler) HandleMessage(conn *websocket.Conn, msg *wsmsg.Packet) {
	if h.Next != nil {
		h.Next.HandleMessage(conn, msg)
	}
}

type AddBillHandler struct {
	Next MessageHandler
}

func (handler *AddBillHandler) HandleMessage(conn *websocket.Conn, msg *wsmsg.Packet) {
	if msg.Type == wsmsg.Type_ADD_BILL {
		fmt.Println("Handling AddBill request:", msg.Content)
		// 处理添加账单消息
	}
}

type DeleteBillHandler struct{}

func (handler *DeleteBillHandler) HandleMessage(conn *websocket.Conn, msg *wsmsg.Packet) {
	if msg.Type == wsmsg.Type_DELETE_BILL {
		fmt.Println("Handling DeleteBill request:", msg.Content)
		// 处理删除账单消息
	}
}
