package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout" // layout is used for layouting widgets.
	"gioui.org/op"     // op is used for recording different operations.
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BoardStyle struct {
}

func (board BoardStyle) Layout(gtx layout.Context) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()

	drawMap(gtx.Ops, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			drawTile(gtx, i, j)
		}
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func drawMap(ops *op.Ops, x int, y int) {
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

func drawTile(gtx layout.Context, roomX int, roomY int) {
	roomSize := gtx.Constraints.Max.X / mapSizeX
	roomPadding := roomSize / 50
	x := roomX * roomSize
	y := roomY * roomSize

	stack := op.Save(gtx.Ops)
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
	stack.Load()
}
