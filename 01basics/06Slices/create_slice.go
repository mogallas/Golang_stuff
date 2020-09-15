package main

import (
	"fmt"
)

func main() {
	a := [5]int{3, 9, 11, 89, 87}
	b := a[1:3]
	fmt.Println(b)
}
