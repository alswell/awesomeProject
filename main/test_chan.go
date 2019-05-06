package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(filepath.Split("/a/b/c/d.txt"))
	//return
	c := make(chan int, 10)

	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(0)
	timer.Stop()
	timerSnap := time.Duration(2)
	go func() {
		for {
			select {
			case t, ok := <-ticker.C:
				fmt.Println("ticker:", t, ok, timerSnap)
				if 5 > timerSnap && timerSnap > 3 {
					fmt.Println("reset timer")
					timer.Reset(time.Second)
					//fmt.Println("stop timer")
					//timer.Stop()
				}
				//if timerSnap > 5 {
				//	fmt.Println("reset timer")
				//	timer.Reset(time.Second)
				//}
			case t, ok := <-timer.C:
				fmt.Println("timer:", t, ok)
				//timer.Reset(time.Second * timerSnap)
			}
			timerSnap++
		}
	}()
	<-c
	return
	c <- 1
	c <- 11
	c <- 111
	for i := range c {
		fmt.Println(i)
		fmt.Println(len(c))
	}
	fmt.Println("end")
	return
	close(c)
	select {
	case i, ok := <-c:
		fmt.Printf("%#v: %#v", i, ok)
	default:
		fmt.Println("default")
	}
	return
}
