package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	/*
			   - start generating the text (goroutine)

			   - init the app
			       - start screen
			       - setup styles


		    - mode welcome-screen
		        - show stats
		        - show welcome message + info. support hotkeys??
		    - mode typing-screen
		        - enable terminal mode that observes all the keys pressed



			   - start the counter when the first key is pressed (goroutine)
			   - the engine
			       - if pressed key is not matched to the first one
			           - add wrong style (red)
			       - else
			           - add correct style (green/blue)
			       - if backspace is pressed
			           - revert the style to the default one
		            - when the end is reached -> show stats


			   - calculate and show stats `Characters per Minute`
			       - consider backspace presses?
			   - terminate the app
			       - show cursor, etc.
			       - exit(0)

	*/

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	s.SetContent(0, 0, 'H', nil, defStyle)
	s.SetContent(1, 0, 'i', nil, defStyle)
	s.SetContent(2, 0, '!', nil, defStyle)

	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventKey:
			_, run := ev.Key(), ev.Rune()
			fmt.Printf("\r%v - %v", run, string(run))
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}
