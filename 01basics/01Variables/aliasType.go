package main

import (
	"fmt"
)

type aliastype = float32

func main() {
	var a aliastype = 10.5
	var b float32 = 10.5
	fmt.Printf("The value of a and b are %v %v and The type of rescpective vars are %T %T\n", a, b, a, b)
	fmt.Println("Is a equal to b?", a == b)
}
