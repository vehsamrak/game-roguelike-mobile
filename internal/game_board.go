package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

const (
	tag          = "controlPressed"
	controlWest  = "west"
	controlSouth = "south"
	controlEast  = "east"
	controlNorth = "north"
)

type ControlsState struct {
	westControlPressed  bool
	eastControlPressed  bool
	southControlPressed bool
	northControlPressed bool
}

type BoardStyle struct {
	controlState *ControlsState
}

func (board BoardStyle) Layout(gtx layout.Context) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()

	drawMap(gtx)
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			drawTile(gtx, i, j)
		}
	}
	board.drawControls(gtx)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func (board BoardStyle) drawControls(gtx layout.Context) {
	defer op.Save(gtx.Ops).Load()

	controlSize := gtx.Constraints.Max.X / 5
	controlSizePadding := gtx.Constraints.Max.X / 20

	westControl := board.drawControl(
		gtx,
		controlWest,
		gtx.Constraints.Min.X,
		gtx.Constraints.Max.Y-controlSize,
		gtx.Constraints.Min.X+controlSize,
		gtx.Constraints.Max.Y,
	)

	southControl := board.drawControl(
		gtx,
		controlSouth,
		westControl.Min.X+controlSize+controlSizePadding,
		westControl.Min.Y,
		westControl.Max.X+controlSize+controlSizePadding,
		westControl.Max.Y,
	)

	_ = board.drawControl(
		gtx,
		controlEast,
		southControl.Min.X+controlSize+controlSizePadding,
		southControl.Min.Y,
		southControl.Max.X+controlSize+controlSizePadding,
		southControl.Max.Y,
	)

	_ = board.drawControl(
		gtx,
		controlNorth,
		southControl.Min.X,
		southControl.Min.Y-controlSize-controlSizePadding,
		southControl.Max.X,
		southControl.Max.Y-controlSize-controlSizePadding,
	)
}

func (board BoardStyle) drawControl(gtx layout.Context, eventTag string, xMin int, yMin int, xMax int, yMax int) clip.Rect {
	defer op.Save(gtx.Ops).Load()

	for _, ev := range gtx.Queue.Events(eventTag) {
		if eventTag == controlWest {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					board.controlState.westControlPressed = true
				case pointer.Release:
					board.controlState.westControlPressed = false
				}
			}
		} else if eventTag == controlEast {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					board.controlState.eastControlPressed = true
				case pointer.Release:
					board.controlState.eastControlPressed = false
				}
			}
		} else if eventTag == controlSouth {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					board.controlState.southControlPressed = true
				case pointer.Release:
					board.controlState.southControlPressed = false
				}
			}
		} else if eventTag == controlNorth {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					board.controlState.northControlPressed = true
				case pointer.Release:
					board.controlState.northControlPressed = false
				}
			}
		}
	}

	controlSizePadding := gtx.Constraints.Max.X / 20

	controlsColor := color.NRGBA{R: 0xFA, G: 0xFA, B: 0xD2, A: 0xFF}
	colorRed := color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}

	op.Offset(f32.Pt(float32(controlSizePadding), -float32(controlSizePadding))).Add(gtx.Ops)
	pointMin := image.Pt(xMin, yMin)
	pointMax := image.Pt(xMax, yMax)

	control := clip.Rect{Min: pointMin, Max: pointMax}
	control.Add(gtx.Ops)
	pointer.Rect(image.Rect(control.Min.X, control.Min.Y, control.Max.X, control.Max.Y)).Add(gtx.Ops)
	pointer.InputOp{Tag: eventTag, Types: pointer.Press | pointer.Release}.Add(gtx.Ops)

	var buttonColor color.NRGBA
	if eventTag == controlWest {
		if board.controlState.westControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlNorth {
		if board.controlState.northControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlSouth {
		if board.controlState.southControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlEast {
		if board.controlState.eastControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else {
		buttonColor = controlsColor
	}

	paint.ColorOp{Color: buttonColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return control
}

func drawMap(gtx layout.Context) {
	defer op.Save(gtx.Ops).Load()

	clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)}.Add(gtx.Ops)
	paint.ColorOp{
		Color: color.NRGBA{
			R: 0x60,
			G: 0x60,
			B: 0x60,
			A: 0xFF,
		},
	}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}

func drawTile(gtx layout.Context, roomX int, roomY int) {
	defer op.Save(gtx.Ops).Load()

	roomSize := gtx.Constraints.Max.X / mapSizeX
	roomPadding := roomSize / 50
	x := roomX * roomSize
	y := roomY * roomSize

	roomSize = roomSize - roomPadding

	op.Offset(f32.Pt(float32(x), float32(y))).Add(gtx.Ops)
	clip.Rect{Min: image.Pt(roomPadding, roomPadding), Max: image.Pt(roomSize, roomSize)}.Add(gtx.Ops)

	var roomColor color.NRGBA
	if roomX == mapSizeX/2 && roomY == mapSizeY/2 {
		roomColor = color.NRGBA{R: 0xFF, G: 0xA5, B: 0x00, A: 0xFF}
	} else {
		roomColor = color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xFF}
	}
	paint.ColorOp{Color: roomColor}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)
}
