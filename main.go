package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type States struct {
	regular string // the initial state. No modificators;
	correct string // color the char into green;
	wrong   string // color the char into red;
	focus   string // add underline symbol to the char to mark as a cursor;
}

// Single character instance
type Char struct {
	value    string // the new value set from the pressed key;
	original string // the initial value taken from the placeholder;
	state    string // chosen style to character;
	toPrint  string
}

// Update the state of reffered character
func (char *Char) setState(val, stt string) {
	char.value = val
	char.state = stt
	char.toPrint = val
}

type Styles struct {
	okBlue, fail, end, underline string
	// HEADER := "\033[95m"
	// OKBLUE = "\033[94m"
	// OKCYAN = "\033[96m"
	// OKGREEN = "\033[92m"
	// WARNING = "\033[93m"
	// FAIL = "\033[91m"
	// ENDC = "\033[0m"
	// BOLD = "\033[1m"
	// UNDERLINE = "\033[4m"
	// BLACK = "\033[97m"
}

// Return a character according its state at the moment of iteration
func (char Char) String() string {
	// TODO: find more elegant solution to store styles
	styles := Styles{
		okBlue:    "\033[94m",
		fail:      "\033[91m",
		end:       "\033[0m",
		underline: "\033[4m",
	}

	var getPrintString = func(style string) string {
		return fmt.Sprintf("%s%s%s", style, char.value, styles.end)
	}
	switch char.state {
	case states.correct:
		char.toPrint = getPrintString(styles.okBlue)
		return getPrintString(styles.okBlue)
	case states.wrong:
		char.toPrint = getPrintString(styles.fail)
		return getPrintString(styles.fail)
	case states.regular:
		char.toPrint = fmt.Sprintf("%s", char.value)
		return fmt.Sprintf("%s", char.value)
	default:
		char.toPrint = char.value
		return char.value
	}
}

const placeholder = "Lorem Ipsum"

var states = States{"r", "c", "w", "f"}

func main() {
	cursor := 0

	var chars []Char

	// TODO: is there any map() alternative from python?
	for _, value := range placeholder {
		chars = append(chars, Char{
			value:    string(value),
			original: string(value),
			state:    states.regular,
		})
	}

	// Handle keyboard input properly
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// timerSeconds := 0

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyBackspace2:
			cursor = max(0, cursor-1)
			chars[cursor].setState(chars[cursor].original, states.regular)
		case keyboard.KeySpace:
			chars[cursor].setState(" ", states.correct)
			cursor += 1
		case keyboard.KeyEsc:
			return
		default:
			if string(char) == chars[cursor].original {
				chars[cursor].setState(string(char), states.correct)
			} else {
				chars[cursor].setState(string(char), states.wrong)
			}
			cursor += 1
		}

		fmt.Printf(RenderText(chars))

		// size := WindowDimensions{}
		// size.getCurrent()

	}

}
