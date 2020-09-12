package main

import (
	"fmt"
)

func calculate(a, b, c, d int) (int, int) {

	total := a + b + c + d
	percentage := (total * 100) / 400
	return total, percentage
}

func main() {
	a, b, c, d := 95, 87, 92, 83
	fmt.Println("Maths, science , computers and english scores are ", a, b, c, d)
	total, percentage := calculate(a, b, c, d)
	fmt.Println("Total score and percentage:", total, percentage)
}
