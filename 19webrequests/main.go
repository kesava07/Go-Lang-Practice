package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://lco.dev")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response %T \n", response)

	fmt.Println(response)

	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)

	content := string(dataByte)

	fmt.Print(content)
}
