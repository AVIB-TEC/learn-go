package main

import "time"

func main() {
	size := 3
	var buffChan = make(chan int, size)

	//reader
	go func() {
		for {
			_ = <-buffChan
			time.Sleep(time.Second * 3)
		}
	}()

	//writer
	writer := func() {
		for i := 1; i <= 10; i++ {
			buffChan <- i
			println(i)
		}
	}

	writer()
}
