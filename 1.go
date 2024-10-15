package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker says: %s...\n", word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	say("nanonets ap automation")
}
