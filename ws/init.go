package ws

import (
	"encoding/json"
	"fmt"
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/gorilla/websocket"
	"time"
)

var LongConnObj LongConn

type LongConn struct {
	Conn *websocket.Conn
	send chan api_struct.Message
}

func (l LongConn) PushSendChanMessage(msg api_struct.Message) {
	l.send <- msg
}

func (l LongConn) wsConn() {
	l.heartbeat()
	l.read()
}

func (l LongConn) heartbeat() {
	// heart beat
	go func() {
		tickerr := time.NewTicker(time.Second * 5)
		defer tickerr.Stop()
		for {
			select {
			case <-tickerr.C:
				if l.Conn != nil {
					l.Conn.WriteMessage(websocket.PingMessage, nil)
				}
			}
		}
	}()
}

func (l LongConn) read() {
	// read message
	go func() {
		for {
			var err error
			if l.Conn == nil {
				l.Conn, _, err = websocket.DefaultDialer.Dial(
					"ws://127.0.0.1:10040?userID=charlie", nil,
				)
				if err != nil {
					panic(err)
				}
				err = l.Conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					panic(err)
				}
				LongConnObj = l
			}
			l.Conn.SetPongHandler(
				func(appData string) error {
					fmt.Println("pong", appData)
					return nil
				},
			)
			messageType, message, err := l.Conn.ReadMessage()
			if err != nil {
				panic(err)
			}
			fmt.Println(messageType, message)
		}
	}()
}

func (l LongConn) write() {
	// write message
	go func() {
		for {
			select {
			case message, ok := <-l.send:
				// send message to server
				if ok {
					err := l.Conn.SetWriteDeadline(
						time.Now().Add(
							time.
								Second * 10,
						),
					)
					if err != nil {

					}
					req, _ := json.Marshal(message.Message)
					l.Conn.WriteMessage(
						websocket.BinaryMessage, req,
					)
				}
			}
		}
	}()
}
