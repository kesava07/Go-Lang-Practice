package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to TIME")
	var presentTime = time.Now()

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.January, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006"))
}
