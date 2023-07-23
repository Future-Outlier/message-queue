// file2.go
package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func receiveMsg() {
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("Received:", msg)
		case data := <-msgChan2: // 修改此行
			newBody := &Person{}
			err := proto.Unmarshal(data, newBody) // 修改此行
			fmt.Println("Received Elliot", err)
			fmt.Println("Received Elliot", newBody.Age, newBody.Name)
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
