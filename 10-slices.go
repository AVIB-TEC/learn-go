package main

import "fmt"

func main() {
	arr1 := [6]int{1, 3, 5, 7, 11, 13}
	slice1 := []int{1, 3, 5, 7, 11, 13}
	slice2 := slice1[1:4]
	var slice3 = make([]int, 2, 3)

	fmt.Println("11- ", arr1)   //	arr1 := [6]int{1, 3, 5, 7, 11, 13}
	fmt.Println("12- ", slice1) // slice1 := []int{1, 3, 5, 7, 11, 13}
	fmt.Println("13- ", slice2) // slice2 := slice1[1:4]
	fmt.Println("14- ", slice3) //	var slice3 = make([]int, 2, 3)

	slice3 = slice1[1:4]

	fmt.Println("18- ", slice3) // 	slice3 = slice1[1:4]
	fmt.Println("19- ", len(slice3))

	slice1 = append(slice1, 200, 300, 400)
	fmt.Println("22- ", slice1) // 	slice1 := []int{1, 3, 5, 7, 11, 13}

	slice2 = append(slice2, []int{7, 8, 9}...)
	fmt.Println("25- ", slice2) //  [3 5 7]

	copyslice := make([]int, len(slice1))
	fmt.Println("28- ", copyslice)
	copy(copyslice, slice1)
	fmt.Println("30- ", copyslice)
}
