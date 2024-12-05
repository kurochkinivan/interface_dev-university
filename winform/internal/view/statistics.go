package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kurochkinivan/winform/internal/controller"
	"github.com/kurochkinivan/winform/internal/model"
)

func NewStatisticsWindow(a *fyne.App, shown *bool, label string) *fyne.Window {
	statisticsWindow := createNewWindow(a, shown, label)
	(*statisticsWindow).SetContent(widget.NewLabel(label))

	statisticsSubscriber := controller.NewStatisticsSubscriber()
	statisticsSubscriber.SetCallBack(func(r []model.Record) {
		updateStatisticsWindow(statisticsWindow, statisticsSubscriber.GetRecords())
	})

	updateStatisticsWindow(statisticsWindow, statisticsSubscriber.GetRecords())

	return statisticsWindow
}

func updateStatisticsWindow(w *fyne.Window, records []model.Record) {
	var items []fyne.CanvasObject
	for _, record := range records {
		text := fmt.Sprintf("Игрок: %s, Счёт: %d, Дата: %v, Уровень: %d",
			record.PlayerName, record.Score, record.DatePlayed.Format("2006-01-02 15:04:05"), record.Level)
		label := widget.NewLabel(text)
		items = append(items, label)
	}

	content := container.NewVBox(items...)
	(*w).SetContent(content)
}
