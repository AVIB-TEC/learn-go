package util

const X int = 100 //exported
const x int = 200 //unexported

// exported
func GetX() int {
	return x
}

// unexported
func getX() int {
	return x
}
