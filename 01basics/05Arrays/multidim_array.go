package main

import "fmt"

func main() {
	a := [4][2]int{{0, 1}, {1, 0}, {0, 0}, {3,2}}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] of array is %.1d\n", i, j, a[i][j])
		}

	}
}
