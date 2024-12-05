package internal

import (
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

		tileXIdx := (cursorX - offsetX) / (tileSize + tileMargin)
		tileYIdx := (cursorY - offsetY) / (tileSize + tileMargin)

		if !(tileXIdx >= 0 && tileXIdx < len(g.board[0]) && tileYIdx >= 0 && tileYIdx < len(g.board)) {
			return nil
		}

		if g.selectedX == tileXIdx && g.selectedY == tileYIdx {
			return nil
		}

		if g.board.CanConnect(g.selectedX, g.selectedY, tileXIdx, tileYIdx) {
			println("SUCCESS!")
			return nil 
		}
		
		g.selectedX = tileXIdx
		g.selectedY = tileYIdx

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
