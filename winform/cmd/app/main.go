package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"github.com/kurochkinivan/winform/internal/constants"
	"github.com/kurochkinivan/winform/internal/model"
	"github.com/kurochkinivan/winform/internal/view"
)

func main() {
	a := app.New()
	w := a.NewWindow("Мое приложение")
	w.Resize(constants.DefaultSize)

	menu := view.CreateMenu(&a)
	w.SetMainMenu(menu)

	w.Show()
	w.SetMaster()
	a.Run()
}

type ExampleSubscriber struct {
	ID       string
	settings *model.Settings
}

// Update - метод, вызываемый при изменении в Settings.
func (es *ExampleSubscriber) Update() {
	fmt.Printf("Subscriber %s notified: PathToStatistics is now '%s'\n", es.ID, es.settings.PathToStatistics)
}
