package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	var slice []int = a[1:4]
	fmt.Println("Array and Slice before:", a, slice)
	for i := range slice {
		slice[i]++

	}
	fmt.Println("Slice after updating", slice)
}
