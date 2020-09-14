package main

import (
	"fmt"
)

func addNewValue(a [4]int) {
	a[0] = 05
	fmt.Println("a value inside function", a)
}

func main() {
	a := [...]int{10, 20, 30, 40}
	fmt.Println("Elements of array before passing arrays as a parameter", a)

	addNewValue(a)
	fmt.Println("Elements of array After passing arrays as a parameter", a)

}
