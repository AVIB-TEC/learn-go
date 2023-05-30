// https://www.digitalocean.com/community/tutorials/importing-packages-in-go
package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

import "strconv"
import m "math"

func main() {
	name := "cloudacademy"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(uuid.New())
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
	fmt.Println(m.Round(f))
}
