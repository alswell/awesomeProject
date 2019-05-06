package main

import (
	. "awesomeProject/pkg"
	"fmt"
	"io/ioutil"
	"strings"
)

type A struct {
	a int
}
type B struct {
	A
	AA
	b int
}
func F(v *[2]int) {
	(*v)[0] = (*v)[1]
	fmt.Println("len:", len(*v), v)
}
func ftest() {
	fmt.Println("ftest")
}


func main() {
	str := "BaC"
	fmt.Printf("%c\n", str[0] - 'A' + 'a')
	x := B{}
	x.A.a = 1
	x.AA.A = 2
	fmt.Printf("%#v", x)

	//resp, err := http.Get("http://www.baidu.com")
	s := strings.NewReader("Hello World!")
	ra, _ := ioutil.ReadAll(s)
	fmt.Printf("%s", ra)
	// Hello World!
}

//func main() {
//	v1 := [...]int{1, 2}
//	F(&v1)
//	fmt.Println("len:", len(v1), v1)
//	//var c chan int
//	c := make(chan int, 10)
//	fmt.Println("hello", len(c), cap(c))
//	c <- 10
//	c <- 20
//	fmt.Println("hello", len(c), cap(c))
//	i := <- c
//	fmt.Println(i, len(c), cap(c))
//}

//func main() {
//	// 构建一个通道
//	ch := make(chan int)
//	// 开启一个并发匿名函数
//	go func() {
//		fmt.Println("start goroutine", len(ch), cap(ch))
//		// 通过通道通知main的goroutine
//		<-ch
//		ch <- 0
//		fmt.Println("start goroutine", len(ch), cap(ch))
//		ch <- 0
//		fmt.Println("start goroutine", len(ch), cap(ch))
//		fmt.Println("exit goroutine")
//	}()
//	fmt.Println("wait goroutine")
//	// 等待匿名goroutine
//	ch <- 10
//	<-ch
//	<-ch
//	fmt.Println("all done")
//	time.Sleep(time.Second * 2)
//}