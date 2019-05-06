package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main()  {
	num := runtime.NumCPU()
	fmt.Println(num)
	runtime.GOMAXPROCS(1)

	s := []int{1, 2, 3, 4, 5}
	for i := range s {
		fmt.Println("i:", i)
		wg.Add(1)
		go func(x *int) {
			fmt.Println(*x)
			wg.Done()
		}(&s[i])
	}
	wg.Wait()
}
