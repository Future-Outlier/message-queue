// file2.go
package main

import (
	"fmt"
	"time"
)

func receiveMsg() {
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("Received:", msg)
		case <-time.After(5 * time.Second): // 超時設置為5秒，可以根據需要修改
			fmt.Println("Timeout! No messages received for 5 seconds.")
			// return
		}
	}
}
