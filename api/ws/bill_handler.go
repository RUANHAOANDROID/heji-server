package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"heji-server/domain"
	"heji-server/pkg"
	"heji-server/wsmsg"
)

type AddBillHandler struct {
	BillUseCase domain.BillUseCase
}

func (h *AddBillHandler) HandleMessage(packet *wsmsg.Packet, ctx context.Context, conn *websocket.Conn) {
	if packet.Type == wsmsg.Type_ADD_BILL {
		fmt.Println("添加账单", packet.Content)
		var bill domain.Bill
		err := json.Unmarshal([]byte(packet.Content), &bill)
		//hex, err := primitive.ObjectIDFromHex(packet.Id)
		if err != nil {
			pkg.Log.Println(err)
		}
		h.BillUseCase.SaveBill(ctx, &bill)
	}
}

type DeleteBillHandler struct {
	BillUseCase domain.BillUseCase
}

func (h *DeleteBillHandler) HandleMessage(packet *wsmsg.Packet, ctx context.Context, conn *websocket.Conn) {
	if packet.Type == wsmsg.Type_DELETE_BILL {
		fmt.Println("删除账单", packet.Content)
	}
}
