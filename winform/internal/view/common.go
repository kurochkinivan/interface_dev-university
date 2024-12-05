package view

import (
	"fyne.io/fyne/v2"
	"github.com/kurochkinivan/winform/internal/constants"
)

func createNewWindow(a *fyne.App, shown *bool, label string) *fyne.Window {
	w := (*a).NewWindow(label)
	*shown = true
	w.Resize(constants.DefaultSize)
	w.Show()
	w.SetOnClosed(func() {
		*shown = false
	})
	return &w
}
