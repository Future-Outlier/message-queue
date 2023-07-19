package main

import "fmt"

type MessageQueue struct {
	queue chan string
}

func (q *MessageQueue) Enqueue(msg string) {
	q.queue <- msg
}

func (q *MessageQueue) Dequeue() string {
	return <-q.queue
}

func main() {
	fmt.Println("Hello, world!")
	str := "Hello, world!"
	q := &MessageQueue{queue: make(chan string, 1)} // Create a buffered channel with capacity 1
	q.Enqueue(str)
	fmt.Println(q.Dequeue())
}
