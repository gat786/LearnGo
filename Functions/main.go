package main

import "fmt"

func main() {
	age := 12
	height := 24

	fmt.Println("Age is ", age, " height is ", height)

	// Formatted strings

	fmt.Printf("my age is %v and my height is %v \n", age, height)

	// quoted strings
	var name = "Ganesh"
	fmt.Printf("my name is %q age is %q and my height is %q \n", name, age, height)

	// print types
	var floatType = 255.5
	fmt.Printf("Float is %v in float format %0.2f and type is %T", floatType, floatType, floatType)
}
