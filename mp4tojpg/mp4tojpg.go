package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/k0kubun/pp"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			pp.Print(r)
			//
			// out, err := os.Create("test.jpg")
			// if err != nil {
			// 	fmt.Printf("error %v", err)
			// 	return
			// }
			//
			// defer out.Close()
			//
			// _, err = io.Copy(out, r.Body)
			// if err != nil {
			// 	fmt.Printf("error %v", err)
			// 	return
			// }

			file, header, err := r.FormFile("file")

			if err != nil {
				fmt.Printf("error %v", err)
				return
			}

			defer file.Close()

			out, err := os.Create(header.Filename)
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

	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error listening:", err)
	}

}
