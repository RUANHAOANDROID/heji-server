package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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
		var protoBill wsmsg.Bill
		for _, anyMsg := range packet.GetContent() {
			// 解析 Any 消息到 Bill 消息
			options := proto.UnmarshalOptions{}
			if err := anypb.UnmarshalTo(anyMsg, &protoBill, options); err != nil {
				fmt.Println("Error unmarshaling Any to Bill:", err)
				return
			}
			// 打印解析后的 Bill 消息
			fmt.Println("Deserialized Bill:", &protoBill)
		}
		bill := protoToBill(&protoBill)
		h.BillUseCase.SaveBill(ctx, &bill)
	}
}
func protoToBill(proto *wsmsg.Bill) domain.Bill {
	hex, err := primitive.ObjectIDFromHex(proto.Id)
	if err != nil {
		pkg.Log.Println(err)
	}
	d := domain.Bill{
		ID:         hex,
		BookId:     proto.BookId,
		Money:      proto.Money,
		Category:   proto.Category,
		CreateUser: proto.CreateUser,
		CreateTime: proto.CreateTime,
		UpdateTime: proto.UpdateTime,
		Remark:     proto.Remark,
		Images:     proto.Images,
	}
	return d
}

type DeleteBillHandler struct {
	BillUseCase domain.BillUseCase
}

func (h *DeleteBillHandler) HandleMessage(packet *wsmsg.Packet, ctx context.Context, conn *websocket.Conn) {
	if packet.Type == wsmsg.Type_DELETE_BILL {
		fmt.Println("删除账单", packet.Content)
	}
}
