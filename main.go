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
}

// Update the state of reffered character
func (char *Char) setState(val, stt string) {
	char.value = val
	char.state = stt
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
		underline: "\033[4m",
		end:       "\033[0m",
	}

	var reprVal string
	switch char.state {
	case styles.okBlue:
		reprVal = fmt.Sprintf("%s%s%s", styles.okBlue, char.value, styles.end)
	case styles.fail:
		reprVal = fmt.Sprintf("%s%s%s", styles.fail, char.value, styles.end)
	default:
		reprVal = char.value
	}
	return reprVal
}

const placeholder = "Lorem Ipsum"

func main() {
	cursor := 0
	states := States{"r", "c", "w", "f"}

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

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyBackspace2 {
			cursor = max(0, cursor-1)
			chars[cursor].setState(chars[cursor].original, states.regular)
		} else {
			if cursor <= len(placeholder) {
				if string(char) == chars[cursor].original {
					chars[cursor].setState(string(char), states.correct)
				} else {
					chars[cursor].setState(string(char), states.wrong)
				}
				cursor += 1
			}
		}

		if cursor >= len(placeholder) || key == keyboard.KeyEsc {
			break
		}

		// NOTE: add styling to test and print to stdout
		fmt.Printf("\r%v", chars)
	}

	// TODO: replace with Stringer interface
	for n := 0; n <= len(placeholder); n++ {
		if cursor == max(cursor, len(placeholder)) {
			break
		}
	}
}
