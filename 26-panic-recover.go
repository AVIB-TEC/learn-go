package main

import "fmt"

func syst() int {
	fmt.Println("system started...")

	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
		fmt.Println(msg)
	}("blah")

	var data []int
	var x = data[0] //causes runtime spanic!
	x++

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := syst()
	fmt.Println(data)

	panic("die!") //exits program with non-zero code
}
