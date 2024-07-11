package main

import (
	"fmt"
	"os"
	"time"
)

type Screen struct {
	done   chan bool
	ticks  chan time.Time
	events chan Event
}

func InitScreen() Screen {
	s := Screen{
		done:   make(chan bool, 1),
		ticks:  make(chan time.Time),
		events: make(chan Event),
	}
	fmt.Print(CLEAN_SCREEN)
	fmt.Print(HIDE_CURSOR)
	fmt.Print(RESET_CURSOR)
	return s
}

func (s *Screen) terminate() {
	close(s.done)
	close(s.ticks)
	close(s.events)

	fmt.Print(CLEAN_SCREEN)
	fmt.Print(SHOW_CURSOR)
	fmt.Print(RESET_CURSOR)

	os.Exit(0)
}
