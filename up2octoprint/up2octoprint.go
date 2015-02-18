package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetURL string) error {
	var bodyBuf bytes.Buffer
	bodyWriter := multipart.NewWriter(&bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//ファイルハンドル操作をオープンする
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Postではヘッダーのセットが出来ない
	// resp, err := http.Post(targetURL, contentType, &bodyBuf)

	client := &http.Client{}
	req, err := http.NewRequest("POST", targetURL, &bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("X-Api-Key", "7BAA6E96DE3B4A34AD31147551CC5A43")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}

// sample usage
func main() {
	//?apikey=7BAA6E96DE3B4A34AD31147551CC5A43
	//targetURL := "http://192.168.11.51:9999/api/files/local?apikey=7BAA6E96DE3B4A34AD31147551CC5A43"
	targetURL := "http://192.168.11.51:9999/api/files/local"
	filename := "./test.gcode"
	postFile(filename, targetURL)
}
