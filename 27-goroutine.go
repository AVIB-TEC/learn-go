package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	n := rand.Intn(3) // n will be between 0 and 3
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomething1(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {
	rand.Seed(time.Now().Unix())

	doSomething1("sync1")

	go doSomething1("async1")
	go doSomething1("async2")
	go doSomething1("async3")

	doSomething1("sync2")

	time.Sleep(time.Second * 10)
}
