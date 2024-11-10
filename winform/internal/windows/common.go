package windows

import (
	"errors"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kurochkinivan/winform/internal/constants"
)

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

func createNewWindow(a *fyne.App, shown *bool, label string) fyne.Window {
	*shown = true
	w := (*a).NewWindow(label)
	w.Resize(constants.DefaultSize)
	w.Show()
	w.SetOnClosed(func() {
		*shown = false
	})
	return w
}
