package main

import (
	"fmt"
	"reflect"
)

func main () {
	//p := new(int)
	//p = nil
	var iface interface{}
	iface = 1
	fmt.Println(iface, reflect.ValueOf(iface))
	fmt.Printf("%#v, %#v", iface, reflect.ValueOf(iface))
}
