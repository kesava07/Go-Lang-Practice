package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Structs in go lang")
	// No Inheritance in GoLang, no Super neither Parent

	chenna := User{"Chenna Kesava", "kesava07.ip@gmail.com", true, 26}

	fmt.Println(chenna)
	fmt.Printf("Chenna details are %+v \n", chenna)

	fmt.Printf("My name is %v and email is %v \n", chenna.Name, chenna.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
