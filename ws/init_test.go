package ws

import (
	"github.com/gorilla/websocket"
	"testing"
)

func TestLongConn_wsConn(t *testing.T) {
	type fields struct {
		Conn *websocket.Conn
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l := LongConn{
					Conn: tt.fields.Conn,
				}
				l.wsConn()
			},
		)
	}
}
