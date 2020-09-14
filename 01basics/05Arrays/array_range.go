//This program is for finding the sum of elements in an array.package 05Arrays
//Author: Srinivasarao Mogallapu
// iterating using range. range returns both index and value

package main

import (
	"fmt"
)

func main() {
	a := [...]int{10, 20, 30, 40}
	sum := 0
	for i, v := range a {
		fmt.Printf("%d the element of Array id %.2d\n", i, v)
		sum += v
	}
	fmt.Println("Total sum of elements if an array is", sum)
}
