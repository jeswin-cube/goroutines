package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	go say(1, "nanonets ap automation")
	go say(2, "nanonets ap automation is awesome")
	go say(3, "nanonets ap automation is awesome")
	go say(4, "nanonets ap automation is awesome")
	go say(5, "nanonets ap automation is awesome")
	go say(6, "nanonets ap automation is awesome")
	time.Sleep(200 * time.Millisecond)
}
