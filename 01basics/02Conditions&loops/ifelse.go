package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("The number is even ")
	} else {
		fmt.Println("The number is odd")
	}

}
