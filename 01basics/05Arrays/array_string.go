package main

import (
	"fmt"
)

func main() {
	a := [...]int{5, 10, 15, 20}
	fmt.Println(a)

	b := [...]string{"USA", "UAE", "India", "England"}
	fmt.Println(b)

	c := b
	c[0] = "France"
	fmt.Println(c)
	fmt.Println("how many countries?", len(c))
}
