package handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	proto2 "heji-server/proto"
)

type AddBillHandler struct{}

func (h *AddBillHandler) HandleMessage(conn *websocket.Conn, b []byte) {

	var addBill proto2.AddBill
	err := proto.Unmarshal(b, &addBill)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(addBill.BillID)
	//save to db
	//转发
	//ack
}

type DeleteBillHandler struct{}

func (h *DeleteBillHandler) HandleMessage(conn *websocket.Conn, b []byte) {
	var deleteBill proto2.DeleteBill
	err := proto.Unmarshal(b, &deleteBill)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(deleteBill)
	//delete db
	//转发
	//ack
}
