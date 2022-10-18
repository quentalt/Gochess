package main

import "fyne.io/fyne/app"

func main() {
	a := app.New()
	w := a.NewWindow("Chessboard")

	grid := createGrid()
	w.SetContent(grid)

}
