package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Maps")

	languages := make(map[string]string)

	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["NG"] = "Angular"

	fmt.Println("List of Languages", languages)

	delete(languages, "JS")
	fmt.Println("List of Languages", languages)

	fmt.Println(languages["NG"])

	for _, value := range languages {
		// fmt.Printf("For key %v, value is %v \n", key, value)
		fmt.Printf("Value is %v \n", value)

	}
}
