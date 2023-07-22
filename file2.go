// file2.go
package main

import (
	"fmt"
)

func receiveMsg() {
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("Received:", msg)
		case <-msgChan2:
			fmt.Println("Received Elliot")
		default:
			// panic("No messages received")
			// fmt.Println("No messages received")
			// case <-time.After(5 * time.Second): // 超時設置為5秒，可以根據需要修改
			// 	fmt.Println("Timeout! No messages received for 5 seconds.")
			// 	// return
			// }
		}
	}
}
