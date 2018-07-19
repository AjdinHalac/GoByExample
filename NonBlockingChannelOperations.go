package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("received a message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <- messages:
		fmt.Println("receieved message", msg)
	case sig := <- signals:
		fmt.Println("receieved a signal", sig)
	default:
		fmt.Println("no activity")
	}
}
