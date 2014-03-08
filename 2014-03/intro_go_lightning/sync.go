package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go say(wg, "let's go!", 3*time.Second)
	go say(wg, "ho!", 2*time.Second)
	go say(wg, "hey!", 1*time.Second)
	wg.Wait()
}

func say(wg *sync.WaitGroup, text string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(text)
	wg.Done()
}

// END OMIT
