package main

import (
	"fmt"
)

func main() {
	array := [4]int{1, 2, 3, 4}
	slice1 := array[:3]
	slice2 := array[1:4]
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 6)

	fmt.Printf("Array = %v , slice = %v\n", array, slice1)
	fmt.Printf("Array = %v , slice = %v\n ", array, slice2)
	fmt.Printf("Length and capacity of slice is %v %v\n", len(slice2), cap(slice2))
}
