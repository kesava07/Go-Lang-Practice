package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to class on Files")

	content := "This is content should need to go to MyFile"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkError(err)

	fmt.Println("File is created successfully and the length is", length)

	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("The content is ", string(dataByte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
