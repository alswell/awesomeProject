package main

import (
	"fmt"
	"sync"
)

func main () {
	ftest()
	lock := sync.RWMutex{}
	lock.RLocker()
	//lock.Lock()
	//fmt.Println(1)
	lock.Unlock()
	fmt.Println(11)
	lock.RLock()
	fmt.Println(2)
	lock.RLock()
	fmt.Println(3)
}
