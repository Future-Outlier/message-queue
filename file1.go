package main

import (
	"time"
)

var msgChan = make(chan string)

func main() {
	go sendMsg()
	receiveMsg() // main 會直接結束，不會等待 goroutine 結束
}

func sendMsg() {
	time.Sleep(2 * time.Second)
	msgChan <- "Hello from file1.go"
}
