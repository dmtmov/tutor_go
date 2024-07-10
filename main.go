package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

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
	CURSOR_X0_Y0 = "\033[H"
	CLEAN_SCREEN = "\033[2J"
	SHOW_CURSOR  = "\033[?25h"
	HIDE_CURSOR  = "\033[?25l"
	NEW_LINE     = "\012"
)

type States struct {
	regular string // the initial state. No modificators;
	correct string // color the char into green;
	wrong   string // color the char into red;
	focus   string // add underline symbol to the char to mark as a cursor;
}

type Sign struct {
	newValue  string // the new value set from the pressed key;
	origValue string // the initial value taken from the placeholder;
	state     string // chosen style to character;
	toPrint   string
}

// Update the state of reffered character
func (s *Sign) setState(val, stt string) {
	s.newValue = val
	s.state = stt
	s.toPrint = val
}

type Event struct {
	r rune
	k keyboard.Key
}

func startKeyObserver(ch chan<- Event) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		runa, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		ch <- Event{runa, key}
	}
}

func startTicking(ch chan<- time.Time, done <-chan bool) {
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			ch <- t
			// fmt.Print("\033[H")
			// fmt.Printf("%s\n", t)
		case <-done:
			return
		}
	}
}

var states = States{"r", "c", "w", "f"}

func (s *Sign) String() string {
	getPrintString := func(style string) string {
		return fmt.Sprintf("%s%s%s", style, s.newValue, ENDC)
	}
	switch s.state {
	case states.correct:
		s.toPrint = getPrintString(OKBLUE)
		return getPrintString(OKBLUE)
	case states.wrong:
		s.toPrint = getPrintString(FAIL)
		return getPrintString(FAIL)
	case states.regular:
		s.toPrint = s.newValue
		return s.newValue
	default:
		s.toPrint = s.newValue
		return s.newValue
	}
}

func RenderText(characters []Sign) string {
	toPrint := make([]string, len(characters))
	for _, val := range characters {
		toPrint = append(toPrint, val.String())
	}

	return fmt.Sprintf("%v", strings.Join(toPrint, ""))
}

func main() {
	// ANSI Escape Codes:
	// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

	tickCh := make(chan time.Time)
	keyCh := make(chan Event)
	done := make(chan bool, 1)
	placeholder := `package main; import "fmt"; func main() { for i := 0; i < 10; i++ { fmt.Println(i) } }`

	cursor := 0
	var chars []Sign
	for _, value := range placeholder {
		chars = append(chars, Sign{
			newValue:  string(value),
			origValue: string(value),
			state:     states.regular,
		})
	}

	go startTicking(tickCh, done)
	go startKeyObserver(keyCh)

	// moves cursor to home position and erase the entire screen
	fmt.Print(CLEAN_SCREEN, CURSOR_X0_Y0)

	for {
		ev := <-keyCh
		fmt.Print(HIDE_CURSOR)

		switch ev.k {
		case keyboard.KeyBackspace2:
			cursor = max(0, cursor-1)
			chars[cursor].setState(chars[cursor].origValue, states.regular)
		case keyboard.KeySpace:
			chars[cursor].setState("_", states.correct)
			cursor += 1
		case keyboard.KeyEsc:
			close(done)
			fmt.Print(SHOW_CURSOR)
			fmt.Print(CLEAN_SCREEN)
			fmt.Print(CURSOR_X0_Y0)
			return
		default:
			if string(ev.r) == chars[cursor].origValue {
				chars[cursor].setState(string(ev.r), states.correct)
			} else {
				chars[cursor].setState(chars[cursor].origValue, states.wrong)
			}
			cursor += 1

			tick := <-tickCh
			fmt.Print(CURSOR_X0_Y0)
			fmt.Print(NEW_LINE)
			fmt.Printf("%s\n", tick)
		}
		fmt.Printf("%s", RenderText(chars))
	}
}
