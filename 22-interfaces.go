package main

import "fmt"

type device interface {
	turnOn() string
}

type iphone struct {
	name  string
	model string
}

type imac struct {
	name  string
	model string
}

func (phone iphone) turnOn() string {
	return "iOS starting up..."
}

func (mac imac) turnOn() string {
	return "macOS starting up..."
}

func main() {
	dev1 := iphone{"iPhone", "11 Pro"}
	dev2 := imac{"iMac", "27 5k Retina"}

	devices := []device{dev1, dev2}

	for _, dev := range devices {
		fmt.Println(dev.turnOn())
	}
}
