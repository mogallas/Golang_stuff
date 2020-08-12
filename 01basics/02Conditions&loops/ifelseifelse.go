package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your Age")
	var age int
	fmt.Scan(&age)
	fmt.Println("your age is", age)
	if age >= 60 {
		fmt.Print("your senior citizon")
	} else if age <= 30 {
		fmt.Println("You are younger")
	} else {
		fmt.Println("you are middle aged person")
	}
}
