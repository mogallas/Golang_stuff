package main

import (
	"cnu_module/interest"
	"fmt"
)

func main() {
	fmt.Println("Simple interest calculation")
	p := 10000.00
	r := 5.0
	t := 2.0
	si := interest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
