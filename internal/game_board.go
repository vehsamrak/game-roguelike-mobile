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

	drawMap(gtx)
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			drawTile(gtx, i, j)
		}
	}
	drawControls(gtx)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func drawControls(gtx layout.Context) {
	defer op.Save(gtx.Ops).Load()

	controlSizePadding := gtx.Constraints.Max.X / 20
	controlSize := gtx.Constraints.Max.X / 5

	controlsColor := color.NRGBA{R: 0xFA, G: 0xFA, B: 0xD2, A: 0xFF}
	// colorRed := color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}

	// west
	state := op.Save(gtx.Ops)
	op.Offset(f32.Pt(float32(controlSizePadding), -float32(controlSizePadding))).Add(gtx.Ops)
	westPointMin := image.Pt(gtx.Constraints.Min.X, gtx.Constraints.Max.Y-controlSize)
	westPointMax := image.Pt(gtx.Constraints.Min.X+controlSize, gtx.Constraints.Max.Y)
	clip.Rect{
		Min: westPointMin,
		Max: westPointMax,
	}.Add(gtx.Ops)
	paint.ColorOp{Color: controlsColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	state.Load()

	// south
	state = op.Save(gtx.Ops)
	op.Offset(f32.Pt(float32(controlSizePadding), -float32(controlSizePadding))).Add(gtx.Ops)
	southPointMin := image.Pt(westPointMin.X+controlSize+controlSizePadding, westPointMin.Y)
	southPointMax := image.Pt(westPointMax.X+controlSize+controlSizePadding, westPointMax.Y)
	clip.Rect{
		Min: southPointMin,
		Max: southPointMax,
	}.Add(gtx.Ops)
	paint.ColorOp{Color: controlsColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	state.Load()

	// east
	state = op.Save(gtx.Ops)
	op.Offset(f32.Pt(float32(controlSizePadding), -float32(controlSizePadding))).Add(gtx.Ops)
	eastPointMin := image.Pt(southPointMin.X+controlSize+controlSizePadding, southPointMin.Y)
	eastPointMax := image.Pt(southPointMax.X+controlSize+controlSizePadding, southPointMax.Y)
	clip.Rect{
		Min: eastPointMin,
		Max: eastPointMax,
	}.Add(gtx.Ops)
	paint.ColorOp{Color: controlsColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	state.Load()

	// north
	state = op.Save(gtx.Ops)
	op.Offset(f32.Pt(float32(controlSizePadding), -float32(controlSizePadding))).Add(gtx.Ops)
	northPointMin := image.Pt(southPointMin.X, southPointMin.Y-controlSize-controlSizePadding)
	northPointMax := image.Pt(southPointMax.X, southPointMax.Y-controlSize-controlSizePadding)
	clip.Rect{
		Min: northPointMin,
		Max: northPointMax,
	}.Add(gtx.Ops)
	paint.ColorOp{Color: controlsColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	state.Load()
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
