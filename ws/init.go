package ws

import (
	"encoding/json"
	"github.com/charlie-bit/go-chat-client/model/api_struct"
	"github.com/gorilla/websocket"
	"time"
)

var LongConnObj *LongConn

type Config struct {
	WSURL string
}

type LongConn struct {
	Config
	Conn *websocket.Conn
	send chan api_struct.Message
}

func (l *LongConn) PushSendChanMessage(msg api_struct.Message) {
	l.send <- msg
}

func (l *LongConn) WsConn() {
	l.initBasicInfo()
	l.heartbeat()
	l.read()
	l.write()
}

func (l *LongConn) initBasicInfo() {
	l.send = make(chan api_struct.Message, 10)
	LongConnObj = l
}

func (l *LongConn) heartbeat() {
	// heart beat
	go func() {
		tickerr := time.NewTicker(time.Second * 5)
		defer tickerr.Stop()
		for {
			select {
			case <-tickerr.C:
				if l.Conn != nil {
					_ = l.Conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
					l.Conn.WriteMessage(websocket.PingMessage, nil)
				}
			}
		}
	}()
}

func (l *LongConn) read() {
	// read message
	go func() {
		for {
			var err error
			if l.Conn == nil {
				l.Conn, _, err = websocket.DefaultDialer.Dial(
					l.Config.WSURL, nil,
				)
				if err != nil {
					panic(err)
				}
				err = l.Conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					panic(err)
				}
				LongConnObj.Conn = l.Conn
			}
			l.Conn.SetPongHandler(
				func(appData string) error {
					return nil
				},
			)
			_, _, err = l.Conn.ReadMessage()
			if err != nil {
				panic(err)
			}
		}
	}()
}

func (l *LongConn) write() {
	// write message
	go func() {
		for {
			if l.Conn == nil {
				continue
			}
			select {
			case message, ok := <-l.send:
				if ok {
					// send message to server
					err := l.Conn.SetWriteDeadline(
						time.Now().Add(
							time.
								Second * 10,
						),
					)
					if err != nil {
						continue
					}
					req, _ := json.Marshal(message.Message)
					err = l.Conn.WriteMessage(
						websocket.BinaryMessage, req,
					)
					if err != nil {
						continue
					}
				}
			}
		}
	}()
}
