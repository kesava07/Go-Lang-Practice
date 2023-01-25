package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello this is to learn GET request")
	performGetRequest()
}

func performGetRequest() {
	var myUrl = "http://localhost:8000"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code ,", response.StatusCode)
	fmt.Println("Content length is ,", response.ContentLength)

	var responseString strings.Builder

	content, err := ioutil.ReadAll(response.Body)
	byteContent, _ := responseString.Write(content)

	fmt.Println(byteContent)

	fmt.Println(responseString.String())
	// fmt.Println(string(content))
	defer response.Body.Close()
}
