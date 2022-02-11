package main

import "fmt"

func main() {
	var name string = "Ganesh Tiwari"
	var nameAgain = "Rahul Bharati"
	var nameThree string

	fmt.Println(name, nameAgain, nameThree)
	nameThree = "bowser"

	fmt.Println(name, nameAgain, nameThree)

	nameFour := "Something other than that"

	fmt.Println(nameFour)

	var age = 1
	ageTwo := 2

	fmt.Println(age, ageTwo)

	var bit8 int8 = 12
	var unsignedBit8 uint8 = 24

	fmt.Println(bit8, unsignedBit8)

	var scoreOne = 25.98
	var scoreTwo float64 = 691623719872380917823078120978

	fmt.Println(scoreOne, scoreTwo)
}
