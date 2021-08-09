package main

import (
	"log"
	"os"

	guiapp "gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"

	"github.com/vehsamrak/game-roguelike-mobile/internal/app"
)

const (
	WindowSizeX = 600
	WindowSizeY = 1000
	mapSizeX    = 11
	mapSizeY    = 11
)

var controlState *app.ControlsState

func main() {
	controlState = &app.ControlsState{}

	go func() {
		window := guiapp.NewWindow(
			guiapp.Title("Roguelike"),
			guiapp.Size(unit.Dp(WindowSizeX), unit.Dp(WindowSizeY)),
		)
		if err := run(window); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	guiapp.Main()
}

func run(window *guiapp.Window) error {
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
	gameBoard := &app.GameBoard{
		ControlState: controlState,
		MapSizeX:     mapSizeX,
		MapSizeY:     mapSizeY,
	}

	return layout.Center.Layout(
		gtx,
		gameBoard.Layout,
	)
}
