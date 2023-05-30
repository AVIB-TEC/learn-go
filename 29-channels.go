package main

import "fmt"

func main() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "Cloud"
		msgChan <- "Academy"
		msgChan <- "2020"
	}()

	msg1 := <-msgChan
	msg2 := <-msgChan
	msg3 := <-msgChan

	fmt.Println(msg1, msg2, msg3)
}
