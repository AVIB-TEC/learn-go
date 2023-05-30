package main

import "fmt"

var x int = 11

func print(id int, x int) {
	fmt.Printf("%d: x=%d\n", id, x)
}

func main() {
	print(1, x)

	x := 2

	print(2, x)

	func(x int) {
		print(3, x)
		if x := 3; x == 3 {
			//x := 100
			print(4, x)
		}
		print(5, x)
	}(5)

	print(6, x)
}
