package main

import (
	"fmt"
)

func main() {
	const a = 10
	fmt.Println("constant value is", a)
	/*a = 20
	fmt.Println("new value of a is", a)
	./constants.go:10:4: cannot assign to a */
}
