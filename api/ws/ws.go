package ws

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"heji-server/pkg"
	"heji-server/wsmsg"
	"time"
)

func serializeMessage(packet *wsmsg.Packet) ([]byte, error) {
	// 序列化消息
	serializedMsg, err := proto.Marshal(packet)
	if err != nil {
		return nil, err
	}
	return serializedMsg, nil
}
func AckOK(id string) ([]byte, error) {
	ackOk := wsmsg.Ack{Msg: "OK", Status: wsmsg.AckStatus_SUCCESS}
	anya, err := anypb.New(&ackOk)
	if err != nil {
		pkg.Log.Println(err)
	}
	ack := wsmsg.Packet{
		Id:        id,
		Type:      wsmsg.Type_ACK,
		Timestamp: time.Now().Unix(),
		Content:   []*anypb.Any{anya},
	}
	return serializeMessage(&ack)
}
func AckErr(id, error string) ([]byte, error) {
	ackOk := wsmsg.Ack{Msg: error, Status: wsmsg.AckStatus_FAILURE}
	anya, err := anypb.New(&ackOk)
	if err != nil {
		pkg.Log.Println(err)
	}
	ack := wsmsg.Packet{
		Id:        id,
		Type:      wsmsg.Type_ACK,
		Timestamp: time.Now().Unix(),
		Content:   []*anypb.Any{anya},
	}
	return serializeMessage(&ack)
}
