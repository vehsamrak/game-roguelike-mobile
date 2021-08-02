package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	windowWidth := unit.Dp(600)
	windowHeight := unit.Dp(1000)

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
	img, _, _, err := loadImageFromFile("/home/petr/git/game-roguelike-mobile/Icon.png")
	if err != nil {
		return err
	}

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
			drawImage(ops, img)
			event.Frame(ops)
			// window.Invalidate()
		}
	}
}

func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4)))
	paint.PaintOp{}.Add(ops)
}

func loadImageFromFile(filePath string) (img image.Image, height int, width int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	img, _, err = image.Decode(file)

	// imageConfig, _, err := image.DecodeConfig(file)
	// if err != nil {
	// 	return nil, 0, 0, err
	// }

	return img, 0, 0, err
}
