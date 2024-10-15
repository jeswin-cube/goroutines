package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func say(wg *sync.WaitGroup, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go say(&wg, 1, "nanonets ap automation")

	wg.Add(1)
	go say(&wg, 2, "nanonets ap automation is awesome")

	wg.Wait()
}
