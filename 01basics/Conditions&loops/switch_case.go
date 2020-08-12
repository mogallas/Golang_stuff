package main

import (
	"fmt"
)

func nodes() {
	fmt.Println("Displaying Nodes list")
}

func pods() {
	fmt.Println("Displaying Pods list")
}

func services() {
	var message string = "Displaying Services list"
	fmt.Println(message)
}

func main() {
	fmt.Println("Please enter your Options:\n------------------------- \n1. Nodes list\n2. Pods list\n3. Services List")
	var option int
	fmt.Scan(&option)
	switch {
	case option == 1:
		nodes()
	case option == 2:
		pods()
	case option == 3:
		services()
	default:
		fmt.Println("Please enter valid option either 1 or 2 or 3")
	}

}
