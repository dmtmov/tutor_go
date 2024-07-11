package main

import (
	"fmt"
	"time"
)

type Ticker struct {
	done chan bool
}

func startTicker(ch chan<- time.Time, done <-chan bool) {
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			ch <- t
			fmt.Print(RESET_CURSOR)
			fmt.Printf("%s", t)
		case <-done:
			return
		}
	}
}
