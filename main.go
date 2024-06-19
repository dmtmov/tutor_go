package main

import (
	// "fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

type appStyle struct {
	style tcell.Style
}

func main() {
	// Generate text here. This is gonna be replaced with ollama version.
	// typingText := "Lorem Ipsum"

	// The application will have two modes:
	// - welcome screen with statistics (if any);
	// - typing screen with the text to fill through
	welcomeMode := appStyle{
		tcell.StyleDefault.Foreground(tcell.ColorTomato).Background(tcell.ColorReset),
	}

	// Initialize screen
	app, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := app.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	app.SetStyle(welcomeMode.style)

    // app.EnableFocus()
    // app.DisableMouse()
	// app.HideCursor()
	// app.EnablePaste()
	app.Clear()

    // TODO: 
    // - get the middle index of welcome message
    // - position welcome message in center of screen
    // - wait for any key pressed
    // - prepare the inner box-block with the text-to-type

	w, h := app.Size()
	// app.SetSize(w, h)

    welcomeRunes := []rune("Press any key to start")

	app.SetContent(w/100*30, h/2, rune('_'), welcomeRunes, welcomeMode.style)

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		app.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// Event loop
	for {
		// Update screen
		app.Show()

		// Poll event
		ev := app.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			app.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				app.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				app.Clear()
			}
		}
	}
}
