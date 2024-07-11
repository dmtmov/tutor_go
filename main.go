package main

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

// ANSI Escape Codes:
// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
const (
	HEADER       = "\033[95m"
	OKBLUE       = "\033[94m"
	OKCYAN       = "\033[96m"
	OKGREEN      = "\033[92m"
	WARNING      = "\033[93m"
	BLACK        = "\033[97m"
	FAIL         = "\033[91m"
	ENDC         = "\033[0m"
	BOLD         = "\033[1m"
	UNDERLINE    = "\033[4m"
	RESET_CURSOR = "\033[H" // set position to x=0, y=0
	CLEAN_SCREEN = "\033[2J"
	SHOW_CURSOR  = "\033[?25h"
	HIDE_CURSOR  = "\033[?25l"
	NEW_LINE     = "\012"
)

var (
	signState = State{"r", "c", "w", "f"}
)

func RenderText(chars []Sign) string {
	toPrint := make([]string, len(chars))
	for _, val := range chars {
		toPrint = append(toPrint, val.String())
	}
	return fmt.Sprintf("%v", strings.Join(toPrint, ""))
}

func main() {
	screen := InitScreen()
	text := NewText(`Lorem ipsum dolor sit amet, ...`)
	cursor := 0

	go startTicker(screen.ticks, screen.done)
	go startEventObserver(screen.events)

	fmt.Print(RESET_CURSOR)
	fmt.Print(RenderText(text.signs))

	for {
		// tick := <-screen.tick
		// fmt.Printf("\r%s", tick)

		ev := <-screen.events
		sign := &text.signs[cursor]
		switch ev.keyCode {
		case keyboard.KeyBackspace2:
			cursor = max(0, cursor-1)
			sign := &text.signs[cursor]
			sign.markRegular(sign.origValue)
		case keyboard.KeySpace:
			sign.markCorrect("_")
			cursor += 1
		case keyboard.KeyEsc:
			screen.terminate()
		default:
			if ev.strValue == sign.origValue {
				sign.markCorrect(ev.strValue)
			} else {
				sign.markWrong(sign.origValue)
			}
			cursor += 1
		}
	}
}
