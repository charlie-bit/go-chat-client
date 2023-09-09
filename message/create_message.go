package message

import (
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/charlie-bit/go-chat-client/pkg/constant"
	"github.com/charlie-bit/go-chat-client/pkg/utils"
)

func initMessageBasicInfo(
	message *api_struct.MsgStruct,
	msgFrom, contentType int32,
	userID string,
) error {
	message.CreateTime = utils.GetCurrentTimestampByMill()
	message.SendTime = message.CreateTime
	message.IsRead = false
	message.Status = constant.MsgStatusSending
	message.SendID = userID
	// message.SenderFaceURL = userInfo.FaceURL
	// message.SenderNickname = userInfo.Nickname
	ClientMsgID := utils.GetMsgID(message.SendID)
	message.ClientMsgID = ClientMsgID
	message.MsgFrom = msgFrom
	message.ContentType = contentType
	return nil
}

func CreateTextMessage(userID, text string) (*api_struct.MsgStruct, error) {
	s := api_struct.MsgStruct{}
	err := initMessageBasicInfo(&s, constant.UserMsgType, constant.Text, userID)
	if err != nil {
		return nil, err
	}
	s.TextElem = &api_struct.TextElem{Content: text}
	return &s, nil
}
