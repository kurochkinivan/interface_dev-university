package view

import (
	"fyne.io/fyne/v2"
)

var (
	settingsShown   bool = false
	statisticsShown bool = false
)

func CreateMenu(a *fyne.App) *fyne.MainMenu {
	gameMenu := fyne.NewMenu("Игра", fyne.NewMenuItem("Новая игра", func() {}))

	statisticsMenu := fyne.NewMenu("Статистика", fyne.NewMenuItem("Посмотреть статистику", func() {
		if !statisticsShown {
			NewStatisticsWindow(a, &statisticsShown, "Статистика")
		}
	}))

	settingsMenu := fyne.NewMenu("Настройки", fyne.NewMenuItem("Перейти в настройки", func() {
		if !settingsShown {
			NewSettingsWindow(a, &settingsShown, "Настройки")
		}
	}))

	menu := fyne.NewMainMenu(gameMenu, settingsMenu, statisticsMenu)

	return menu
}
