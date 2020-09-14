package main

import (
	"fmt"
)

func main() {
	a := [...]int{10, 20, 30, 40, 50}

	for i := 0; i < len(a); i++ {
		//fmt.Printf("%d th element of array a is %.2f\n", i, a[i])
		fmt.Printf("%d th element of array a is %.2d\n", i, a[i])
	}
}
