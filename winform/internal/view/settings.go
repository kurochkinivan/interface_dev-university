package view

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kurochkinivan/winform/internal/model"
)

func NewSettingsWindow(a *fyne.App, shown *bool, label string) {
	settingsWindow := createNewWindow(a, shown, label)

	settings := model.GetSettings()

	path := binding.NewString()
	path.Set(settings.PathToStatistics)
	entry := widget.NewEntryWithData(path)

	form := &widget.Form{
		Orientation: widget.Horizontal,
		Items:       []*widget.FormItem{{Text: "Путь к файлу:", Widget: entry}},
		SubmitText:  "Сохранить",
		OnSubmit: func() {
			ok, err := checkFilePath(settingsWindow, entry.Text)
			if err != nil {
				dialog.ShowError(err, *settingsWindow)
				return
			}

			if ok {
				fmt.Println("обновился path", entry.Text)
				settings.UpdatePath(entry.Text)
				path.Set(settings.PathToStatistics)
				dialog.ShowInformation("Успех!", "Данные были успешно обновлены", *settingsWindow)
			}
		},
	}
	(*settingsWindow).SetContent(form)
}

func checkFilePath(w *fyne.Window, fileName string) (bool, error) {
	wd, err := os.Getwd()
	if err != nil {
		return false, err
	}

	wd = filepath.Dir(filepath.Dir(wd))
	_, err = os.Open(filepath.Join(wd, "data", fileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			dialog.ShowInformation("Ошибка", "Файла с указанным именем не существует", *w)
			return false, nil
		}
		return false, err
	}

	return true, nil
}
