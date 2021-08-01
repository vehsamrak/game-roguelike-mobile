package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("OK", func() {
			hello.SetText("")
		}),
		widget.NewButton("Cancel", func() {
			hello.SetText("Cancelled!")
		}),
	))

	w.ShowAndRun()
}
