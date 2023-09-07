package ws

import (
	"fmt"
	"testing"
	"time"
)

func TestTickter(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 定时器触发时执行的逻辑
			fmt.Println("定时器触发")
		}
	}
}
