package api_struct

type MsgStruct struct {
	ClientMsgID          string    `json:"clientMsgID,omitempty"`
	ServerMsgID          string    `json:"serverMsgID,omitempty"`
	CreateTime           int64     `json:"createTime"`
	SendTime             int64     `json:"sendTime"`
	SessionType          int32     `json:"sessionType"`
	SendID               string    `json:"sendID,omitempty"`
	RecvID               string    `json:"recvID,omitempty"`
	MsgFrom              int32     `json:"msgFrom"`
	ContentType          int32     `json:"contentType"`
	SenderNickname       string    `json:"senderNickname,omitempty"`
	SenderFaceURL        string    `json:"senderFaceUrl,omitempty"`
	GroupID              string    `json:"groupID,omitempty"`
	Content              string    `json:"content,omitempty"`
	Seq                  int64     `json:"seq"`
	IsRead               bool      `json:"isRead"`
	Status               int32     `json:"status"`
	IsReact              bool      `json:"isReact,omitempty"`
	IsExternalExtensions bool      `json:"isExternalExtensions,omitempty"`
	AttachedInfo         string    `json:"attachedInfo,omitempty"`
	Ex                   string    `json:"ex,omitempty"`
	LocalEx              string    `json:"localEx,omitempty"`
	TextElem             *TextElem `json:"textElem,omitempty"`
}

type TextElem struct {
	Content string `json:"content"`
}

type GeneralWsReq struct {
	ReqIdentifier int    `json:"reqIdentifier"`
	Token         string `json:"token"`
	SendID        string `json:"sendID"`
	OperationID   string `json:"operationID"`
	MsgIncr       string `json:"msgIncr"`
	Data          []byte `json:"data"`
}

type Message struct {
	Message GeneralWsReq
	Resp    chan *GeneralWsResp
}

type GeneralWsResp struct {
	ReqIdentifier int    `json:"reqIdentifier"`
	ErrCode       int    `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	MsgIncr       string `json:"msgIncr"`
	OperationID   string `json:"operationID"`
	Data          []byte `json:"data"`
}
