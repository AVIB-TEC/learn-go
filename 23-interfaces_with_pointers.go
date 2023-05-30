package main

import "fmt"
import "strings"

type device1 interface {
	turnOn1() string
	update(version float32)
}

type iphone1 struct {
	name    string
	model   string
	version float32
}

type imac1 struct {
	name    string
	model   string
	version float32
}

func (phone iphone1) turnOn1() string {
	return "iOS starting up..."
}

func (mac imac1) turnOn1() string {
	return "macOS starting up..."
}

func (phone *iphone1) update(version float32) {
	phone.version = version
}

func (mac *imac1) update(version float32) {
	mac.version = version
}

func main() {
	dev1 := iphone1{"iPhone", "11 Pro", 13.1}
	dev2 := imac1{"iMac", "27 5k Retina", 10.15}

	devices := []device1{&dev1, &dev2}

	for _, dev := range devices {
		if strings.Contains(dev.turnOn1(), "iOS") {
			dev.update(14.0)
		} else if strings.Contains(dev.turnOn1(), "macOS") {
			dev.update(11.00)
		}
	}

	fmt.Println(dev1)
	fmt.Println(dev2)
}
