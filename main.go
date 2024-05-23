package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type States struct {
	regular, correct, wrong, focus string
}

/*
Describe atomic character in a text to type.

	`value` contains the pressed key value for current iteration;
	`original` stores the initial symbol in a text;
	`state` can have one of States value;
*/
type Char struct {
	value, original, state string
}

// Update the state of reffered character
func (char *Char) setState(val, stt string) string {
	return val + stt
}

// Return a character according its state at the moment of iteration
func (char *Char) repr() string {
	return char.value
}

const placeholder = "Lorem Ipsum"

func main() {
	cursor := 0
	states := States{"r", "c", "w", "f"}

	var charSeq []Char

	for _, value := range placeholder {
		ch := Char{
			value:    string(value),
			original: string(value),
			state:    states.regular,
		}
		charSeq = append(charSeq, ch)
	}

	// Handle keyboard input properly
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// infinite loop
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		cursorChar := &charSeq[cursor]
		if key == keyboard.KeyBackspace {
			cursor = max(0, cursor-1)
			cursorChar.setState(cursorChar.original, states.regular)
			// cursorChar.state = states.regular
			charSeq[cursor].value = charSeq[cursor].original
		} else {
			if cursor <= len(placeholder) {
				if string(char) == charSeq[cursor].original {
					charSeq[cursor].state = states.correct
				} else {
					charSeq[cursor].state = states.wrong
				}
				charSeq[cursor].value = string(char)
				cursor += 1
			}
		}

		// NOTE: add styling to test and print to stdout
		var printSeq []string
		for _, value := range charSeq {
			printSeq = append(printSeq, value.repr())
		}
		fmt.Printf("\r%v", printSeq)

		if cursor >= len(placeholder) || key == keyboard.KeyEsc {
			break
		}
	}

	for n := 0; n <= len(placeholder); n++ {

		fmt.Println(charSeq)
		// textSeq[n] = Char{placeholder[cursor]}
		// fmt.Printf("%d\n", len(placeholder))
		// fmt.Printf("%v", string(placeholder[cursor]))
		cursor += 1

		if cursor == max(cursor, len(placeholder)) {
			fmt.Println()
			break
		}
	}
}
