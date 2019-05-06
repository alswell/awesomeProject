package main

import "fmt"

var intSlice = make([]int, 0)
func cpSlice() (r []int) {
	r = intSlice
	fmt.Printf("cpSlice: %p %p\n", intSlice, r)
	return
}
type SomeInfo struct {
	id int
	name string
}
var someInfo = SomeInfo{1, "zn"}
func cpStruct() (i bool, r *SomeInfo) {
	r = &someInfo
	return
}

func main() {
	i, someInfo2 := cpStruct()
	fmt.Println(i)
	i, _ = cpStruct()
	fmt.Println(i)
	someInfo.id = 2
	fmt.Println(someInfo, someInfo2)
	return
	intSlice = append(intSlice, 11, 22, 33)
	intSlice2 := make([]int, 0)
	intSlice2 = intSlice
	fmt.Printf("1. %p\n", intSlice2)
	fmt.Printf("1. %p\n", intSlice)
	intSlice = append(intSlice, 44, 55)
	des := cpSlice()
	fmt.Printf("%p\n", intSlice2)
	fmt.Printf("%p\n", intSlice)
	fmt.Printf("des %p\n", des)
	fmt.Println(des)
	fmt.Println(intSlice, intSlice2)
	intSlice[1] = 111
	fmt.Printf("%#v\n", des)
	fmt.Println(intSlice, intSlice2)
	//src := make([]int, 0)
	//des := make([]int, 2)
	////var des []int
	//fmt.Printf("%#v\n", des)
	//fmt.Println(src == nil)
	//src = append(src, 11, 22, 33)
	//fmt.Println(copy(des, src))
	//fmt.Println(src[2])
	//fmt.Println(des == nil)
	//src[2] = 222
	//fmt.Println(src[2])
	//fmt.Println(des[2])
}
