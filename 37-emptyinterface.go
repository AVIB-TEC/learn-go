package main

import "fmt"

type company struct {
	name string
}

func main() {
	var a, b, c, d interface{}

	a = 42
	b = "blah"
	c = true
	d = company{"cloudadacademy"}

	func(list ...interface{}) {
		for _, v := range list {
			fmt.Printf("%v, %T\n", v, v)
		}

	}(a, b, c, d)
}
