package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to Maths in GoLang")

	var numberOne int = 2
	var numberTwo float64 = 4

	fmt.Println("The sum of two numbers: ", numberOne+int(numberTwo))

	// Random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// Random number generation with Crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
