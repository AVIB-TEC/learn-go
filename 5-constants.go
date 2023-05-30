package main

import (
	"fmt"
)

//const maxValue //compile error

const population int = 10000

func main() {
	const name = "CloudAcademy"

	//name = "Blah" //compile error

	if true {
		const color = "red"

		fmt.Println(population)
		fmt.Println(name)
	}

	//fmt.Println(color) //compile error
}
