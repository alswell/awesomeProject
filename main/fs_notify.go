package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"strings"
	"sync"
)

var wgFsNotify = sync.WaitGroup{}
func watch(filename string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	//defer watcher.Close()

	wgFsNotify.Add(1)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("<-watcher.Events fail", ok)
					return
				}
				log.Println(filename, "event:", event)
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("remove file:", event.Name)
					if strings.Contains(filename, "itomagent.conf") {
						fmt.Println("add itomagent.conf")
						err = watcher.Add(filename)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println("<-watcher.Errors fail", ok)
					return
				}
				log.Println(filename, "error:", err)
			}
		}
	}()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	watch("/home/zhouning/go/src/metric_monitor/itom_agent/conf")
	watch("/home/zhouning/go/src/metric_monitor/itom_agent/conf/itomagent.conf")
	wgFsNotify.Wait()
	//done := make(chan bool)
	//<-done
}
