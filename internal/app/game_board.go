package app

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

type GameBoard struct {
	ControlState *ControlsState
	GameMap      *GameMap
	MapSizeX     int
	MapSizeY     int
}

func (gb *GameBoard) Layout(gtx layout.Context) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()

	gb.drawMap(gtx)
	gb.drawControls(gtx)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func (gb GameBoard) drawControls(gtx layout.Context) {
	defer op.Save(gtx.Ops).Load()

	controlSize := gtx.Constraints.Max.X / 5
	controlSizePadding := gtx.Constraints.Max.X / 20

	westControl := gb.drawControl(
		gtx,
		controlWest,
		gtx.Constraints.Min.X,
		gtx.Constraints.Max.Y-controlSize,
		gtx.Constraints.Min.X+controlSize,
		gtx.Constraints.Max.Y,
	)

	southControl := gb.drawControl(
		gtx,
		controlSouth,
		westControl.Min.X+controlSize+controlSizePadding,
		westControl.Min.Y,
		westControl.Max.X+controlSize+controlSizePadding,
		westControl.Max.Y,
	)

	_ = gb.drawControl(
		gtx,
		controlEast,
		southControl.Min.X+controlSize+controlSizePadding,
		southControl.Min.Y,
		southControl.Max.X+controlSize+controlSizePadding,
		southControl.Max.Y,
	)

	_ = gb.drawControl(
		gtx,
		controlNorth,
		southControl.Min.X,
		southControl.Min.Y-controlSize-controlSizePadding,
		southControl.Max.X,
		southControl.Max.Y-controlSize-controlSizePadding,
	)
}

func (gb GameBoard) drawControl(
	gtx layout.Context,
	eventTag string,
	xMin int,
	yMin int,
	xMax int,
	yMax int,
) clip.Rect {
	defer op.Save(gtx.Ops).Load()

	for _, ev := range gtx.Queue.Events(eventTag) {
		if eventTag == controlWest {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					gb.ControlState.westControlPressed = true
				case pointer.Release:
					gb.ControlState.westControlPressed = false
				}
			}
		} else if eventTag == controlEast {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					gb.ControlState.eastControlPressed = true
				case pointer.Release:
					gb.ControlState.eastControlPressed = false
				}
			}
		} else if eventTag == controlSouth {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					gb.ControlState.southControlPressed = true
				case pointer.Release:
					gb.ControlState.southControlPressed = false
				}
			}
		} else if eventTag == controlNorth {
			if x, ok := ev.(pointer.Event); ok {
				switch x.Type {
				case pointer.Press:
					gb.ControlState.northControlPressed = true
				case pointer.Release:
					gb.ControlState.northControlPressed = false
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
		if gb.ControlState.westControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlNorth {
		if gb.ControlState.northControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlSouth {
		if gb.ControlState.southControlPressed {
			buttonColor = colorRed
		} else {
			buttonColor = controlsColor
		}
	} else if eventTag == controlEast {
		if gb.ControlState.eastControlPressed {
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

func (gb *GameBoard) drawMap(gtx layout.Context) {
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

	for i := 0; i < gb.MapSizeX; i++ {
		for j := 0; j < gb.MapSizeY; j++ {
			gb.drawTile(gtx, i, j)
		}
	}
}

func (gb *GameBoard) drawTile(gtx layout.Context, roomX int, roomY int) {
	defer op.Save(gtx.Ops).Load()

	roomSize := gtx.Constraints.Max.X / gb.MapSizeX
	roomPadding := roomSize / 50
	x := roomX * roomSize
	y := roomY * roomSize

	roomSize = roomSize - roomPadding

	op.Offset(f32.Pt(float32(x), float32(y))).Add(gtx.Ops)
	clip.Rect{Min: image.Pt(roomPadding, roomPadding), Max: image.Pt(roomSize, roomSize)}.Add(gtx.Ops)

	paint.ColorOp{Color: gb.tileColor(roomX, roomY)}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)
}

func (gb *GameBoard) tileColor(roomX int, roomY int) (roomColor color.NRGBA) {
	// character tile
	characterX, characterY := gb.GameMap.character.XY()
	if roomX == characterX && roomY == characterY {
		return color.NRGBA{R: 0xDC, G: 0x14, B: 0x3C, A: 0xFF}
	}

	roomTile := gb.GameMap.FindTileByXY(roomX, roomY)
	if roomTile == nil {
		return color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xFF}
	}

	// color by tile type
	switch roomTile.Type {
	case TileTypeForest:
		// limegreen
		roomColor = color.NRGBA{R: 0x32, G: 0xCD, B: 0x32, A: 0xFF}
	case TileTypeWater:
		// lightskyblue
		roomColor = color.NRGBA{R: 0x87, G: 0xCE, B: 0xFA, A: 0xFF}
	case TileTypeMountain:
		// goldenrod
		roomColor = color.NRGBA{R: 0xDA, G: 0xA5, B: 0x20, A: 0xFF}
	case TileTypeCliff:
		// slategray
		roomColor = color.NRGBA{R: 0x70, G: 0x80, B: 0x90, A: 0xFF}
	default:
		roomColor = color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xFF}
	}

	return roomColor
}
