package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	coursesList := []course{
		{"React JS", 300, "react.com", "abc123", []string{"JS", "React"}},
		{"Angular JS", 200, "angular.com", "bcd123", []string{"JS", "Angular"}},
		{"VUE JS", 150, "vue.com", "gdf123", nil},
	}

	finalJson, _ := json.MarshalIndent(coursesList, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React JS",
		"Price": 300,
		"website": "react.com",
		"tags": ["JS","React"]
}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSON is not Valid")
	}

	// some cases where you just want to add data to key values

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and it's type is %T \n", k, v, v)
	}

}
