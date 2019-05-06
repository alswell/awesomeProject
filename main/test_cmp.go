package main

import (
	"fmt"
	"reflect"
)

type Struct4CMP2 struct {
	a, b int
	c, d string
}
type Struct4CMP struct {
	id int
	name string
	course []string
	child Struct4CMP2
}
func main() {
	s1 := Struct4CMP{}
	s2 := Struct4CMP{id: 0}
	fmt.Println(reflect.DeepEqual(&s1, s2))
}
