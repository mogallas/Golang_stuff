package main

import "fmt"

func addOne(slice2 []int) {
	for i := range slice2 {
		slice2[i] += 2
	}
}
func main() {
	slice := []int{10, 20, 30}
	fmt.Println("Slice before passing to function \n", slice)
	addOne(slice)
	fmt.Println("Slice after passing to function \n", slice)
}
