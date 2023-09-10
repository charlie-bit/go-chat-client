package message

import (
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/charlie-bit/go-chat-client/pkg/proto/msg"
	"github.com/charlie-bit/go-chat-client/ws"
	"reflect"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		s       *msg.MsgData
		sendID  string
		recvID  string
		groupID string
	}
	tests := []struct {
		name    string
		args    args
		want    *api_struct.MsgStruct
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test send message",
			args: args{
				sendID:  "charlie1",
				groupID: "charlie_group",
			},
			wantErr: false,
		},
	}
	l := &ws.LongConn{
		Conn: nil,
		Config: ws.Config{
			WSURL: "ws://127.0.0.1:10040?userID=charlie1",
		},
	}
	l.WsConn()
	time.Sleep(time.Second * 3)
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.args.s, _ = CreateTextMessage(tt.args.sendID, "hello, I am charlie")
				got, err := SendMessage(
					tt.args.s, tt.args.sendID, tt.args.recvID, tt.args.groupID,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf(
						"SendMessage() error = %v, wantErr %v", err, tt.wantErr,
					)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("SendMessage() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
	select {}
}
