package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello this is to learn POST request")
	// performPostRequest()
	performPostFormRequest()
}
func performPostFormRequest() {
	myUrl := "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstName", "Chennakesava")
	data.Add("email", "kesava07.ip@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	defer response.Body.Close()
}

func performPostRequest() {
	myUrl := "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name": "Chennakesava",
			"age": 26,
			"gender": "Male"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	defer response.Body.Close()
}
