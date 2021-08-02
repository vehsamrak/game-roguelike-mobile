package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

const (
	GUISizeX       = 600
	GUISizeY       = 1000
	mapSizeX       = 12
	mapSizeY       = 12
	roomSize       = 48
	roomSizeOffset = 2
)

func main() {
	windowWidth := unit.Sp(GUISizeX)
	windowHeight := unit.Sp(GUISizeY)

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
			ops := new(op.Ops)
			drawMap(ops, mapSizeY, mapSizeX)
			for i := 0; i < mapSizeY; i++ {
				for j := 0; j < mapSizeY; j++ {
					drawTile(ops, i, j)
				}
			}
			event.Frame(ops)
		}
	}
}

func drawMap(ops *op.Ops, x int, y int) {
	x = x * (roomSize + roomSizeOffset)
	y = y * (roomSize + roomSizeOffset)

	stack := op.Save(ops)
	clip.Rect{Max: image.Pt(x, y)}.Add(ops)
	paint.ColorOp{
		Color: color.NRGBA{
			R: 0x60,
			G: 0x60,
			B: 0x60,
			A: 0xFF,
		},
	}.Add(ops)
	paint.PaintOp{}.Add(ops)
	stack.Load()
}

func drawTile(ops *op.Ops, x int, y int) {
	x = x * (roomSize + 2)
	y = y * (roomSize + 2)

	stack := op.Save(ops)
	op.Offset(f32.Pt(float32(x), float32(y))).Add(ops)
	clip.Rect{Max: image.Pt(roomSize, roomSize)}.Add(ops)
	paint.ColorOp{
		Color: color.NRGBA{
			R: 0x80,
			G: 0x80,
			B: 0x80,
			A: 0xFF,
		},
	}.Add(ops)
	paint.PaintOp{}.Add(ops)
	stack.Load()
}
