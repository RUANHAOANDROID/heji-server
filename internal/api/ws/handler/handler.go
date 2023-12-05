package handler

import (
	"github.com/gorilla/websocket"
)

const (
	AddBill    = 1
	DeleteBill = 2
)

type MessageHandler interface {
	HandleMessage(conn *websocket.Conn, b []byte)
}

func CreateHandler(msgType int) MessageHandler {
	switch msgType {
	case AddBill:
		return &AddBillHandler{}
	case DeleteBill:
		return &DeleteBillHandler{}
	}
	return nil
}
