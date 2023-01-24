package main

import "fmt"

func main() {
	greeting()
	fmt.Println("Welcome to Functions in Golang")

	adderResult := adder(2, 3)
	fmt.Println("Adder result is ", adderResult)

	proAdderResult, proMessage := proAdder(1, 2, 3, 4, 5, 6, 7)

	fmt.Println("Pro adder result is ", proAdderResult)
	fmt.Println("Pro adder message is ", proMessage)

}

func proAdder(values ...int) (int, string) {
	result := 0

	for _, value := range values {
		result += value
	}
	return result, "Message from Pro Adder"
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}
func greeting() {
	fmt.Println("Namasthey from GoLang")
}
