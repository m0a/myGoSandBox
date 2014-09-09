package main

import (
	"fmt"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloWorld")
}

func name(rw http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.Replace(path, "/", " ", -1)
		fmt.Fprintf(w, "Hello! %s", path)

	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()
		name := queryParam.Get("name")
		age := queryParam.Get("age")
		fmt.Printf("%#v\n\n", r)
		fmt.Printf("%#v\n", r.URL)

		fmt.Fprintf(w, "Hello! %s,your age is %s", name, age)
	})

	http.ListenAndServe(":4000", nil)

}
