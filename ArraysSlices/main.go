package main

import "fmt"

func main() {
	var ages [3]int = [3]int{15, 12, 20}
	fmt.Println(ages)

	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	fmt.Println(names)

	// slices
	var scores = []int{1, 1, 2, 3, 4, 5}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores)

	// ranges
	rangeOne := names[1:3]
	fmt.Println(rangeOne)
}
