package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Arrays")

	var fruitList [4]string

	fruitList[0] = "Banana"
	fruitList[1] = "Apple"
	fruitList[3] = "Peach"

	fmt.Println("Fruits list ", fruitList)
	fmt.Println("Fruits list length ", len(fruitList))

	var veggiesList = [5]string{"Tomato", "Potato", "Mushroom"}
	fmt.Println("Veggies list ", veggiesList)
	fmt.Println("Veggies list length ", len(veggiesList))

}
