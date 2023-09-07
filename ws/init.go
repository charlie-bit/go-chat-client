package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type LongConn struct {
	Conn *websocket.Conn
}

func (l LongConn) wsConn() {
	// heart beat
	go func() {
		tickerr := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-tickerr.C:
				if l.Conn != nil {
					l.Conn.WriteMessage(websocket.PingMessage, nil)
				}
			}
		}
	}()
	// read message
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
}
