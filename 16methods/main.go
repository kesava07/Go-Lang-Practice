package main

import "fmt"

func main() {

	john := User{"John Doe", "john@gmail.com", true, 26}

	fmt.Printf("Hello %v welcome and your statues is %v \n", john.Name, john.Status)

	john.GetStatus()

	john.GetEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Your status is", u.Status)
}

func (u User) GetEmail() {
	u.Email = "doe@gmail.com"
	fmt.Println("Your new email is ", u.Email)
}
