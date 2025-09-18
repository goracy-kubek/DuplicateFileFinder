package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Duplicate finder")
	w.Resize(fyne.NewSize(600, 400))

	w.SetContent(container.NewVBox(
		searchField(),
		widget.NewSeparator(),
	))
	w.ShowAndRun()
}

func searchField() *fyne.Container {
	search := widget.NewEntry()
	search.SetPlaceHolder("Start folder path...")

	button := widget.NewButton("Search", func() {
		println("Searching")
	})

	return container.NewBorder(nil, nil, nil, button, search)
}
