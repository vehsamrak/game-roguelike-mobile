package main

import (
	"fmt"
	"log"
	"math/rand"
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
	WindowSizeX          = 600
	WindowSizeY          = 1000
	mapSizeX             = 11
	mapSizeY             = 11
	mapMaxSizeXMin       = -10
	mapMaxSizeXMax       = 10
	mapMaxSizeYMin       = -10
	mapMaxSizeYMax       = 10
	gameRandomSeed       = 0
	characterStartPointX = 0
	characterStartPointY = 0
)

func main() {
	rand.Seed(gameRandomSeed)

	go func() {
		if mapSizeX%2 == 0 {
			log.Println(fmt.Sprintf("Map size X must always be even. Now it is odd - %d", mapSizeX))
			os.Exit(1)
		}

		window := guiapp.NewWindow(
			guiapp.Title("Roguelike"),
			guiapp.Size(unit.Dp(WindowSizeX), unit.Dp(WindowSizeY)),
		)

		character := app.NewCharacter(characterStartPointX, characterStartPointY)
		gameMap := app.NewGameMap(
			mapMaxSizeXMin,
			mapMaxSizeXMax,
			mapMaxSizeYMin,
			mapMaxSizeYMax,
			character,
		)

		err := run(
			window,
			&app.ControlsState{},
			gameMap,
			app.NewCharacterActionProvider(character, gameMap),
		)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	guiapp.Main()
}

func run(
	window *guiapp.Window,
	controlState *app.ControlsState,
	gameMap *app.GameMap,
	characterActionProvider *app.CharacterActionProvider,
) error {
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

			Layout(gtx, controlState, gameMap, characterActionProvider)

			event.Frame(gtx.Ops)
		}
	}
}

func Layout(
	gtx layout.Context,
	controlState *app.ControlsState,
	gameMap *app.GameMap,
	characterActionProvider *app.CharacterActionProvider,
) layout.Dimensions {
	gameBoard := &app.GameBoard{
		ControlState:            controlState,
		CharacterActionProvider: characterActionProvider,
		GameMap:                 gameMap,
		BoardSizeX:              mapSizeX,
		BoardSizeY:              mapSizeY,
	}

	return layout.Center.Layout(
		gtx,
		gameBoard.Layout,
	)
}
