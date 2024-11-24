package menu

import (
	"fyne.io/fyne/v2"
	"github.com/kurochkinivan/winform/internal/windows"
)

var (
	gameWindowShown       bool = false
	settingsWindowShown   bool = false
	statisticsWindowShown bool = false

	statisticsWindow *fyne.Window

	pathToFile string = "default.json"
)

func CreateMenu(a *fyne.App) *fyne.MainMenu {
	gameMenu := fyne.NewMenu("Игра", fyne.NewMenuItem("Новая игра", func() {
		if !gameWindowShown {
			windows.NewGameWindow(a, &gameWindowShown, "Игра")
		}
	}))

	statisticsMenu := fyne.NewMenu("Статистика", fyne.NewMenuItem("Посмотреть статистику", func() {
		if !statisticsWindowShown {
			statisticsWindow = windows.NewStatisticsWindow(a, &statisticsWindowShown, pathToFile, "Статистика")
		}
	}))

	settingsMenu := fyne.NewMenu("Настройки", fyne.NewMenuItem("Перейти в настройки", func() {
		if !settingsWindowShown {
			windows.NewSettingsWindow(a, &settingsWindowShown, &pathToFile, "Настройки", statisticsWindow)
		}
	}))

	menu := fyne.NewMainMenu(gameMenu, settingsMenu, statisticsMenu)

	return menu
}
