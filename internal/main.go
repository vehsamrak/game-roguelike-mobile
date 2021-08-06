package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

const (
	WindowSizeX = 600
	WindowSizeY = 1000
	mapSizeX    = 11
	mapSizeY    = 11
)

var controlState *ControlsState

func main() {
	controlState = &ControlsState{}

	go func() {
		window := app.NewWindow(
			app.Title("Roguelike"),
			app.Size(unit.Dp(WindowSizeX), unit.Dp(WindowSizeY)),
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
		case system.FrameEvent:
			gtx := layout.NewContext(&op.Ops{}, event)

			Layout(gtx)

			event.Frame(gtx.Ops)
		}
	}
}

func Layout(gtx layout.Context) layout.Dimensions {
	return layout.Center.Layout(
		gtx,
		BoardStyle{
			controlState: controlState,
		}.Layout,
	)
}
