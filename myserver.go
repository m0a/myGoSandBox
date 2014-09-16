package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	ID   int
	Name string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloWorld")
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		//ファイル名を{id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.Replace(path, "/", " ", -1)
		fmt.Fprintf(w, "Hello! %s", path)
	})

	http.HandleFunc("/persons", PersonHandler)

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
