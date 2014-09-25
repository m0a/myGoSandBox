package extchan

import (
	"net/http"
	"code.google.com/p/go.net/websocket"
	"reflect"
)


type unknownTypeChan interface{}

type pairChannel struct {
	fromClient unknownTypeChan
	toClient   unknownTypeChan
}
type channelTable struct {

	map[string]pairChannel
}

var defaultChannelTable channelTable=channelTable{}

func HandleFunc()func(w http.ResponseWriter, req *http.Request){
	return defaultChannelTable.HandleFunc()
}

func MakeChannel(id string,
	fromClientStruct interface{},
	toClientStruct interface {})(fromClient unknownTypeChan,toclient unknownTypeChan) {

	fromClient=reflect.MakeChan(fromClientStruct,1)
	toclient=reflect.MakeChan(toClientStruct,1)


}
func (c *channelTable) HandleFunc()func(w http.ResponseWriter, req *http.Request) {

	handler:=func(ws *websocket.Conn){
		//		ws.JSON.re
//		for {
//			select  {
//			case value:=<-fromClient:
//				//				websocket.JSON.Send(ws,value)
//				//				ws.Read(msg)
//				data,err:=easyjson.NewEasyJson(value)
//				if err!=nil {
//					panic("error")
//				}
//				fmt.Fprintf(ws,"%s",data)
//
//			default:
//				//				JSON.Receive(ws,&data)
//				// msg := make([]byte, 512)
//				readdata,err:=easyjson.NewEasyJson(ws)
//				if err!=nil {
//					panic("error")
//				}
//
//			toClient<-readdata
//
//			}
//		}
	}

	return func (w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(handler)}
	s.ServeHTTP(w, req)

}
}
