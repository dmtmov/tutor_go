package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type appMode struct {
	style tcell.Style
	is_on bool
}

func main() {
	// Generate text here. This is gonna be replaced with ollama version.
	// typingText := "Lorem Ipsum"

	logFilePath := "/tmp/go_tutorial.log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	defer logFile.Close()
	// The application will have two modes:
	// - welcome screen with statistics (if any);
	// - typing screen with the text to fill through
	welcomeMode := appMode{
		tcell.StyleDefault.Foreground(
			tcell.ColorTomato).Background(tcell.ColorReset),
		true,
	}

	// Initialize the application
	app, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := app.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	app.SetStyle(welcomeMode.style)
	app.DisableMouse()
	app.Clear()

	// Position welcome text in the center of screen
	w, h := app.Size()
	welcomeText := "ress any key to start"
	x_, y_ := w/2-len(welcomeText)/2, h/2

	app.SetContent(x_, y_, rune('P'), []rune(welcomeText), welcomeMode.style)

	// TODO:
	// - get the middle index of welcome message
	// - position welcome message in center of screen
	// - wait for any key pressed
	// - prepare the inner box-block with the text-to-type

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

		// we are in the welcome mode now
		log.Println(welcomeMode, x_,  y_)

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
