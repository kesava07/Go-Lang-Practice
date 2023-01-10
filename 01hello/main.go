package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var title string = "Hello world"
	comment := "Hello chenna"
	fmt.Println(title)
	fmt.Println(comment)

	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of Pizza")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating", input)
	fmt.Printf("Rating is of type %T \n", input)
}
