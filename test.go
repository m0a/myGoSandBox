package main

import (
	"fmt"
	"reflect"
)

type AAA struct {
	index int
	value string
}
func call(i interface {}){
	reflect.TypeOf(i).Kind()=reflect.Chan
	a:=reflect.New(typeA)

	fmt.Printf("<%T>",a)
	fmt.Printf("<%T?",i)
}

func call2(ch interface {}){
	fmt.Printf("%T",ch)
}
func main() {

	data:=AAA{1,"str"}

	fmt.Printf("%v",data)
	call(data)

	ch:=make(chan AAA,1)
	call2(ch)

}
