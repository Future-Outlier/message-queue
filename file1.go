package main

import (
	"time"

	proto "github.com/golang/protobuf/proto"
)

var msgChan = make(chan string)
var msgChan2 = make(chan proto.Message)

func main() {
	go sendMsg()
	receiveMsg() // main 會直接結束，不會等待 goroutine 結束
}

func sendMsg() {
	time.Sleep(2 * time.Second)
	msg := "Hello from file1.go"
	msgChan <- msg

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}
	msgChan2 <- elliot
}
