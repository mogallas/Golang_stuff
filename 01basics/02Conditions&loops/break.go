package main

import "fmt"

func main() {
	a := 1
	for a <= 6 {
		fmt.Println("the number is ", a)
		if a == 3 {
			break
		}
		a++
	}
}
