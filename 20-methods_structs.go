package main

import "fmt"

type person3 struct {
	firstname string
	surname   string
	age       int
}

func (p *person3) fullname() string {
	return p.firstname + " " + p.surname
}

func (p *person3) canDrive() bool {
	if p.age >= 20 {
		return true
	} else {
		return false
	}
}

func (p *person3) updateAge(newAge int) {
	p.age = newAge
}

func main() {
	person1 := person3{"John", "Doe", 40}
	person2 := person3{"Mark", "Doe", 19}

	fmt.Printf("%s can drive: %t\n", person1.fullname(), person1.canDrive())
	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())

	person2.updateAge(person2.age + 1) //marks birthday
	fmt.Println(person2.age)

	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())
}
