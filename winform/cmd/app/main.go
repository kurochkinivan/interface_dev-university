package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/kurochkinivan/winform/internal/constants"
	"github.com/kurochkinivan/winform/internal/menu"
)

func main() {
	a := app.New()
	w := a.NewWindow("Мое приложение")
	w.Resize(constants.DefaultSize)

	menu := menu.CreateMenu(&a)
	w.SetMainMenu(menu)

	w.Show()
	w.SetMaster()
	a.Run()
}
