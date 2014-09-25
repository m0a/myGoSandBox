package main

import (
	"fmt"
	"strings"
)


type JsObj interface{}
type JsArr []JsObj
type JsDic  map[string]JsObj

func NewJson(i interface{})JsObj{
	return "1"
}

func (ja JsArr) String()string{
	ret:=make([]string,len(ja))
	for i,v:=range ja {
		switch cv:=v.(type){
		case string:
			ret[i]="\""+cv+"\""
		default:
			ret[i]=fmt.Sprint(cv)
		}
	}
	return "["+strings.Join(ret,",")+"]"
}

func (jd JsDic) String()string{
	var ret []string
	for k,v:=range jd {
		switch cv:=v.(type){
		case string:
			ret=append(ret,"\""+k+"\":\""+cv+"\"")
		default:
			ret=append(ret,"\""+k+"\":"+fmt.Sprint(cv))
		}
	}
	return "{"+strings.Join(ret,",")+"}"
}


func main(){
//	fmt.Println("Hello.")
	jsondata:=JsDic{"a":12,
		"b":1,
		"c":JsArr{"a","b",1,2,3,4},
		"d":JsDic{"a":"a","b":1,"c":12,"d":23.5} }
	fmt.Printf("%v\n",jsondata)

	fmt.Printf("%v\n",jsondata["d"].(JsDic)["c"])


}
