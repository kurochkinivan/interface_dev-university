package internal

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 980
)

type Game struct {
	board     Board
	selectedX int
	selectedY int
}

func NewGame(board Board) *Game {
	return &Game{
		board:     board,
		selectedX: -1,
		selectedY: -1,
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()

		// Вычисляем смещение для центра поля
		offsetX, offsetY := CenterOffSet(ScreenWidth, ScreenHeight)

		tileX := (cursorX - offsetX) / (tileSize + tileMargin)
		tileY := (cursorY - offsetY) / (tileSize + tileMargin)

		if !(0 <= tileX && tileX < len(g.board[0]) && 0 <= tileY && tileY < len(g.board)) {
			return nil
		}

		if g.selectedX == tileX && g.selectedY == tileY {
			return nil
		}

		fmt.Println("selected:", g.selectedX, g.selectedY, "new:", tileX, tileY)
		if g.board.CanConnect(g.selectedX, g.selectedY, tileX, tileY) {
			println("SUCCESS!")
			return nil
		}

		g.selectedX = tileX
		g.selectedY = tileY

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	offsetX, offsetY := CenterOffSet(ScreenWidth, ScreenHeight)
	screen.Fill(color.RGBA{50, 50, 50, 255})

	for _, row := range g.board {
		for _, tile := range row {
			tile.DrawTile(screen, offsetX, offsetY, g.selectedX, g.selectedY)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
