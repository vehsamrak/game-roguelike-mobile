package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/unit"
)

func main() {
	windowWidth := unit.Dp(100)
	windowHeight := unit.Dp(100)

	go func() {
		window := app.NewWindow(
			app.Title("Roguelike"),
			app.Size(windowWidth, windowHeight),
		)
		if err := run(window); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	// var ops op.Ops
	for {
		windowEvent := <-window.Events()
		switch event := windowEvent.(type) {
		case system.DestroyEvent:
			return event.Err
		case key.Event:
			switch event.Name {
			// close window when ESC pressed
			case key.NameEscape:
				return nil
			}
		}
	}
}
