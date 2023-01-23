package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers in Go")

	// var myPointer *int
	// fmt.Println("Value of the my pointer is ", myPointer)

	myPointer := 23

	var ptr = &myPointer

	fmt.Println("Value of the pointer is ", ptr)
	fmt.Println("Value of the pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value of the pointer is ", myPointer)

}
