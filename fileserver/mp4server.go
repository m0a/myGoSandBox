package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("run Server")

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	list, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	//fmt.Printf("list count = %d\n", len(list))

	fmt.Fprintf(w, "<html>")
	for _, finfo := range list {
		if finfo.IsDir() || strings.Index(finfo.Name(), ".mp4") == -1 {
			//strings.Index(finfo.Name(), ".mp4") == -1
			continue
		}

		fmt.Fprintf(w, "a= <a href='/files/%s' >%s</a><br>", finfo.Name(), finfo.Name())
	}
	fmt.Fprintf(w, "</html>")

}
