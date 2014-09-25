package main

import (
	"net/http"
	"fmt"
	"code.google.com/p/go.net/websocket"
	"github.com/m0a/easyjson"
)

//func MakeChanHandleFunc

func channelHandleFunc(
	fromClient chan interface{},
	toClient   chan interface{}) func(w http.ResponseWriter, req *http.Request) {

	echoHandler:=func(ws *websocket.Conn){
//		ws.JSON.re
		for {
			select  {
			case value:=<-fromClient:
//				websocket.JSON.Send(ws,value)
//				ws.Read(msg)
				data,err:=easyjson.NewEasyJson(value)
				if err!=nil {
					panic("error")
				}
				fmt.Fprintf(ws,"%s",data)

			default:
//				JSON.Receive(ws,&data)
				// msg := make([]byte, 512)
				readdata,err:=easyjson.NewEasyJson(ws)
				if err!=nil {
					panic("error")
				}

				toClient<-readdata

			}
		}
	}

	return func (w http.ResponseWriter, req *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(echoHandler)}
	s.ServeHTTP(w, req)

	}
}


func main(){
	fmt.Println("Hello...")

	readch:=make(chan interface{},1)
	writech:=make(chan interface{},1)
	http.HandleFunc("/echo",channelHandleFunc(readch,writech))
	go func(){
		for {
			select {
			case v := <-writech:
				fmt.Println(v)
			readch<-fmt.Sprintf("%s", v)
			}
		}
	}()
	http.ListenAndServe(":12345",nil)

}
