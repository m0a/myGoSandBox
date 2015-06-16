package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
}
