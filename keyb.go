package main

import (
	"github.com/eiannone/keyboard"
)

type Event struct {
	rune     rune
	keyCode  keyboard.Key
	strValue string
}

func NewEvent(runa rune, key keyboard.Key) Event {
	e := Event{}
	e.rune = runa
	e.keyCode = key
	e.strValue = string(runa)
	return e
}

func startEventObserver(ch chan<- Event) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		runa, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		ch <- NewEvent(runa, key)
	}
}
