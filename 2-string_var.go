package main

import (
	"fmt"
)

var (
	Name1 string = "global1"
	Name2 string = "global2"
)

var var1, var2 string = "local1", "local2"

func main() {
	var name1 string
	var name2 string = "CloudAcademy"
	name3 := "DevOps 2020"
	name4 := name2 + " " + name3

	fmt.Println(name1) //prints empty string
	name1 = "Blah"
	fmt.Println(name1) //prints Blah

	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

	fmt.Printf("%v -- %v\n", name2, name3)

	fmt.Println(Name1, Name2, var1, var2)
}
