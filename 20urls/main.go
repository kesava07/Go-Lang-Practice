package main

import (
	"fmt"
	"net/url"
)

var myUrl = "https://lco.dev:3000/courses?course=reactjs&paymentId=b3723bucw723"

func main() {
	fmt.Println("Welcome to class on URLS")

	response, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Scheme)
	fmt.Println(response.Port())
	fmt.Println(response.Path)
	fmt.Println(response.Host)
	fmt.Println(response.RawQuery)

	qPerams := response.Query()

	fmt.Printf("The type of query params %T \n", qPerams)
	fmt.Println(qPerams["course"])

	for key, value := range qPerams {
		fmt.Printf("Key is %v and value is %v \n", key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=chenna",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
