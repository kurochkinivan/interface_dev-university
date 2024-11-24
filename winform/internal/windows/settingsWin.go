package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewSettingsWindow(a *fyne.App, shown *bool, pathToFile *string, label string, statisticsWindow *fyne.Window) {
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
				
				ShowStatisticsData(*pathToFile, statisticsWindow)
			}
		},
	}
	w.SetContent(form)

}
