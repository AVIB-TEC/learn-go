package main

import "fmt"

type company struct {
	name string
}

func main() {
	var x interface{} = company{"CloudAcademy"}

	c1 := x.(company)
	fmt.Println(c1)

	if c2, ok := x.(company); ok {
		fmt.Println(c2, ok)
	}

	//n := x.(int); n++ //runtime panic
}
