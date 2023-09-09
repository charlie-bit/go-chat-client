package message

import (
	"encoding/json"
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/charlie-bit/go-chat-client/pkg/constant"
	"github.com/charlie-bit/go-chat-client/ws"
)

func SendMessage(
	s *api_struct.MsgStruct, sendID, recvID,
	groupID string,
) (*api_struct.MsgStruct, error) {
	data, err := json.Marshal(s)
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
