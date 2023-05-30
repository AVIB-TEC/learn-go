package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	squares := make(chan int, 20)

	for i := f(); i <= 100; i = f() {
		squares <- i
	}

	close(squares)

	for elem := range squares {
		fmt.Println(elem)
	}

}
