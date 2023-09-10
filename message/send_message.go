package message

import (
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/charlie-bit/go-chat-client/pkg/constant"
	"github.com/charlie-bit/go-chat-client/pkg/proto/msg"
	"github.com/charlie-bit/go-chat-client/ws"
	"google.golang.org/protobuf/proto"
)

func SendMessage(
	s *msg.MsgData, sendID, recvID,
	groupID string,
) (*api_struct.MsgStruct, error) {
	s.GroupID = groupID
	data, err := proto.Marshal(s)
	if err != nil {
		return nil, err
	}
	msg := api_struct.Message{
		Message: api_struct.GeneralWsReq{
			ReqIdentifier: constant.SendMsg,
			SendID:        sendID,
			Data:          data,
		},
		Resp: make(chan *api_struct.GeneralWsResp, 1),
	}
	ws.LongConnObj.PushSendChanMessage(msg)
	return nil, nil
}
