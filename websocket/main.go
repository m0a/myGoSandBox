package main

import (
	"code.google.com/p/go.net/websocket"
	_ "fmt"
	_ "io"
	"log"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {

	type T struct {
		Msg   string
		Count int
	}

	// receive JSON type T
	var data T
	websocket.JSON.Receive(ws, &data)

	log.Printf("data=%#v\n", data)

	// send JSON type T
	websocket.JSON.Send(ws, data)

	// var message string
	// websocket.Message.Receive(ws, &message)
	// websocket.Message.Send(ws, message)

	// io.Copy(ws, ws)
	// log.Printf("%#v...\n", ws.Config().Origin)

	// msg := make([]byte, 512)
	// n, err := ws.Read(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Recive: %s\n", msg[:n])

	// _, err = fmt.Fprintf(ws, "change_message:%s", msg[:n])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Send:%s\n", msg[:n])

}

func main() {

	// http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/echo",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(echoHandler)}
			s.ServeHTTP(w, req)
		})

	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
