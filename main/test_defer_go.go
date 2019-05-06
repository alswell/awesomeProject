package main

import (
	"fmt"
	"os"
	"sync"
)

func main () {
	wg := sync.WaitGroup{}
	defer func() { fmt.Println("defer")}()
	go func() {
		defer func() { fmt.Println("defer in go") }()

		wg.Done()
	}()
	wg.Add(1)
	wg.Wait()
	os.Exit(0)
}
