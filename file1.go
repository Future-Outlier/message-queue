package main

import (
	"fmt"
	"time"

	proto "github.com/golang/protobuf/proto"
)

var msgChan = make(chan string)
var msgChan2 = make(chan []byte)

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
	data, err := proto.Marshal(elliot)
	if err == nil {
		msgChan2 <- data
	}
	fmt.Printf("type of a is %T\n", *elliot)

	var m proto.Message
	m = elliot
	if _, ok := m.(*Person); ok {
		fmt.Println(" ok elliot is of type proto.Message")
	} else {
		fmt.Println("elliot is NOT of type proto.Message")
	}

	if proto.MessageName(elliot) != "" {
		fmt.Println(" MessageName elliot is of type proto.Message")
	} else {
		fmt.Println("elliot is NOT of type proto.Message")
	}

}
