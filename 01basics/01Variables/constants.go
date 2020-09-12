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
	//declare multiple constants
	const (
		name = "srini"
		age  = 31
	)
	fmt.Println("name and age is", name, age)

}
