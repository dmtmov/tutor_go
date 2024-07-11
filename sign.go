package main

import (
	"fmt"
)

type State struct {
	regular string // the initial state. No modificators;
	correct string // color the char into green;
	wrong   string // color the char into red;
	focus   string // add underline symbol to the char to mark as a cursor;
}

type Text struct {
	cursor int // indicates the focused position on array of signs
	signs  []Sign
}

type Sign struct {
	state     string // chosen style to character;
	newValue  string // the new value set from the pressed key;
	origValue string // the initial value taken from the placeholder;

	// NOTE: this field should not exist
	styledValue string // contains color codes
}

// Update the state of reffered character
func (s *Sign) setState(val, stt string) {
	s.newValue = val
	s.state = stt
	s.styledValue = val
}

func (s *Sign) markWrong(val string) {
	s.setState(val, signState.wrong)
}

func (s *Sign) markCorrect(val string) {
	s.setState(val, signState.correct)
}

func (s *Sign) markRegular(val string) {
	s.setState(val, signState.regular)
}

// Styled representation of character
func (s *Sign) String() string {
	getPrintString := func(style string) string {
		return fmt.Sprintf("%s%s%s", style, s.newValue, ENDC)
	}

	switch s.state {
	case signState.correct:
		s.styledValue = getPrintString(OKBLUE)
		return getPrintString(OKBLUE)
	case signState.wrong:
		s.styledValue = getPrintString(FAIL)
		return getPrintString(FAIL)
	case signState.regular:
		s.styledValue = s.newValue
		return s.newValue
	default:
		s.styledValue = s.newValue
		return s.newValue
	}
}

func NewText(text string) Text {
	var chars []Sign
	for _, value := range text {
		chars = append(chars, Sign{
			newValue:  string(value),
			origValue: string(value),
			state:     signState.regular,
		})
	}

	return Text{
		cursor: 0,
		signs:  chars,
	}
}
