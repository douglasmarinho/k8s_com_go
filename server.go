package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	nome := os.Getenv("NAME")
	idade := os.Getenv("AGE")

	fmt.Fprint(w, "Hello, I'm %s. I'm %s.", nome, idade)
}
