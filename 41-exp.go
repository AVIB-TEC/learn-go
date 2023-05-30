package main

import (
	"fmt"
	"github.com/antgonto/util"
)

func main() {
	fmt.Println(util.X)
	//fmt.Println(util.x) //compile failure - unexported

	fmt.Println(util.GetX())
	//fmt.Println(util.getX()) //compile failure - unexported

}
