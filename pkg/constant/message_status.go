package constant

// Message status
const (
	MsgStatusDefault int32 = iota
	MsgStatusSending
	MsgStatusSendSuccess
	MsgStatusSendFailed
	MsgStatusHasDeleted
	MsgStatusFiltered
)
