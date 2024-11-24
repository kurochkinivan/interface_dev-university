package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewGameWindow(a *fyne.App, shown *bool, label string) {
	w := createNewWindow(a, shown, label)

	w.SetContent(widget.NewLabel(label))
}
