package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(600, 600))

	// w.SetContent(widget.NewLabel("Hello World!"))
	w.SetContent(widget.NewButton("Click", func() {
		w.SetContent(widget.NewLabel("Hello World!"))
	}))

	w.ShowAndRun()
}
