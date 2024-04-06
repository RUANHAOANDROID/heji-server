package ws

import (
	"google.golang.org/protobuf/proto"
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
	ack := wsmsg.Packet{
		Id:        id,
		Type:      wsmsg.Type_ACK_OK,
		Timestamp: time.Now().Unix(),
		Content:   "",
	}
	return serializeMessage(&ack)
}
func AckErr(id, error string) ([]byte, error) {
	ack := wsmsg.Packet{
		Id:        id,
		Type:      wsmsg.Type_ACK_ERR,
		Timestamp: time.Now().Unix(),
		Content:   error,
	}
	return serializeMessage(&ack)
}
