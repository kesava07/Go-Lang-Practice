package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to MOD in GoLang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello users welcome to MOD topic")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GoLang series in YouTube </h1>"))
}
