package main

import (
	"fmt"
)

type mytype float32

func main() {
	var a mytype = 10.5
	var b float32 = 10.5
	fmt.Printf("The value of a and b are %v %v and The type of rescpective vars are %T %T\n", a, b, a, b)
	//fmt.Println("Is a equals to be?", a == b)
	//./derivedType.go:13:38: invalid operation: a == b (mismatched types mytype and float32)
}
