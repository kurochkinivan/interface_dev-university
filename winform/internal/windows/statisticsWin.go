package windows

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	gamerecords "github.com/kurochkinivan/winform/internal/gameRecords"
)

func NewStatisticsWindow(a *fyne.App, shown *bool, pathToFile, label string) *fyne.Window {
	w := createNewWindow(a, shown, label)
	w.SetContent(widget.NewLabel(label))

	ShowStatisticsData(pathToFile, &w)

	return &w
}

func ShowStatisticsData(pathToFile string, w *fyne.Window) {
	records, err := gamerecords.ReadDataFile(pathToFile)
	if err != nil {
		dialog.ShowError(err, (*w))
		return
	}

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
