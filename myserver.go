package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloWorld")
}

func name(rw http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)

}
