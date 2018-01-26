package main

import (
	"fmt"
	"time"
)

func say(s string, j int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d / %d: %s\n", i, j, s)
		j += 1
	}
}

func main() {
	j := 1
	go say("hello", j)
	say("world", j)
}
