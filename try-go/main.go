package main

import (
	"fmt"

	something "github.com/google/uuid"

	"net/http"
)

func main() {
	greet()

	str := generateUuid()
	fmt.Println(str)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, oke!"))
	})

	http.ListenAndServe(":80", nil)
}

func greet() {
	fmt.Println("Hello, world!")
}

func generateUuid() string {
	return something.NewString()
}
