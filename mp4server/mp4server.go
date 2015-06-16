package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("run Server")

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/thumbnail", thumbnailHandler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
}

// カレントフォルダ内のmp4ファイルの一覧表示
func rootHandler(w http.ResponseWriter, r *http.Request) {

	list, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "<html>")
	for _, finfo := range list {
		if finfo.IsDir() || strings.Index(finfo.Name(), ".mp4") == -1 {
			continue
		}

		fmt.Fprintf(w, "a= <a href='/files/%s' >%s</a><br>", finfo.Name(), finfo.Name())
	}

	str := `
	<form action="http://localhost:9999/thumbnail" method="post" enctype="multipart/form-data">
	<label for="file">Filename:</label>
	<input type="file" name="file" id="file">
	<input type="submit" name="submit" value="Submit">
	</form>
	`

	fmt.Fprintf(w, "%s", str)
	fmt.Fprintf(w, "</html>")

}

//
func thumbnailHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		file, header, err := r.FormFile("file")

		if err != nil {
			fmt.Printf("error %v", err)
			return
		}

		defer file.Close()

		out, err := os.Create("test.jpg")
		if err != nil {
			fmt.Printf("error %v", err)
			return
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("error %v", err)
			return
		}

		fmt.Printf("file upload success! %v", header.Filename)

	} else {
		fmt.Printf("Method is %v", r.Method)
	}
}