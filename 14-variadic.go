package main

import "fmt"

func displayCount(id int, letters ...string) {
	count := 0
	for range letters {
		count++
	}

	//display id, letters count, letters type, and letters content
	fmt.Printf("%d - %d - %T - %s\n", id, count, letters, letters)
}

func main() {
	displayCount(1, "c", "l", "o", "u", "d")
	displayCount(2, "a", "c", "a", "d", "e", "m", "y")

	cloud := []string{"d", "e", "v", "o", "p", "s"}
	displayCount(3, cloud...)
}
