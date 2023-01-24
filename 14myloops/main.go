package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v, and the day is %v \n", index, day)
	// }

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 4 {
			goto lco
		}
		if roughValue == 5 {
			roughValue++
			continue
		}
		fmt.Println("Value is", roughValue)
		roughValue++
	}
lco:
	fmt.Println("Jumping to GoLang")
}
