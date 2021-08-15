package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"

	guiapp "gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"

	"github.com/vehsamrak/game-roguelike-mobile/assets"
	"github.com/vehsamrak/game-roguelike-mobile/internal/app"
)

const (
	WindowSizeX          = 550
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
			app.NewControlsState(),
			gameMap,
			app.NewCharacterActionProvider(character, gameMap),
			app.NewImageMap(assets.FileSystem),
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
	imageMap map[string]image.Image,
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

			Layout(gtx, controlState, gameMap, characterActionProvider, imageMap)

			event.Frame(gtx.Ops)
		}
	}
}

func Layout(
	gtx layout.Context,
	controlState *app.ControlsState,
	gameMap *app.GameMap,
	characterActionProvider *app.CharacterActionProvider,
	imageMap map[string]image.Image,
) layout.Dimensions {
	gameBoard := &app.GameBoard{
		ControlState:            controlState,
		CharacterActionProvider: characterActionProvider,
		GameMap:                 gameMap,
		ImagesMap:               imageMap,
		BoardSizeX:              mapSizeX,
		BoardSizeY:              mapSizeY,
	}

	return layout.Center.Layout(
		gtx,
		gameBoard.Layout,
	)
}
