package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on Slices")

	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Println("Fruits list", fruitList)
	fmt.Printf("Type of Fruits list %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Papaya")

	fmt.Println("New Fruits list", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 235
	highScores[1] = 765
	highScores[2] = 453
	highScores[3] = 140
	// highScores[4] = 543

	highScores = append(highScores, 333, 777, 456)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"React JS", "Javascript", "Angular", "Node JS", "Java", "Go Lang"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
