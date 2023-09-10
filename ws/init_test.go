package ws

import (
	"github.com/gorilla/websocket"
	"testing"
)

func TestLongConn_wsConn(t *testing.T) {
	type fields struct {
		Conn *websocket.Conn
		Cfg  Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test ws conn",
			fields: fields{
				Conn: nil,
				Cfg: Config{
					WSURL: "ws://127.0.0.1:10040?userID=charlieinit",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l := &LongConn{
					Conn:   tt.fields.Conn,
					Config: tt.fields.Cfg,
				}
				l.WsConn()
			},
		)
	}
	select {}
}
