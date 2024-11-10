package windows

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	gamerecords "github.com/kurochkinivan/winform/internal/gameRecords"
)

func NewStatisticsWindow(a *fyne.App, shown *bool, pathToFile, label string) {
	w := createNewWindow(a, shown, label)
	w.SetContent(widget.NewLabel(label))

	records, err := gamerecords.ReadDataFile(pathToFile)
	if err != nil {
		dialog.ShowError(err, w)
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

	w.SetContent(content)

}

func NewGameWindow(a *fyne.App, shown *bool, label string) {
	w := createNewWindow(a, shown, label)

	w.SetContent(widget.NewLabel(label))
}

func NewSettingsWindow(a *fyne.App, shown *bool, pathToFile *string, label string) {
	w := createNewWindow(a, shown, label)

	path := binding.NewString()
	path.Set(*pathToFile)
	entry := widget.NewEntryWithData(path)

	form := &widget.Form{
		Orientation: widget.Horizontal,
		Items:       []*widget.FormItem{{Text: "Путь к файлу:", Widget: entry}},
		SubmitText:  "Сохранить",
		OnSubmit: func() {
			ok, err := checkFilePath(&w, entry.Text)
			if err != nil {
				dialog.ShowError(err, w)
			}
			if ok {
				*pathToFile = entry.Text
				path.Set(*pathToFile)
				dialog.ShowInformation("Успех!", "Данные были успешно обновлены", w)
			}
		},
	}
	w.SetContent(form)
}
