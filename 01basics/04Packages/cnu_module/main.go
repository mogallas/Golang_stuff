package main

import (
	"cnu_module/interest"
	"fmt"
)

func main() {
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := interest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
