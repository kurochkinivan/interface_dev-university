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
	selectedXAxe int
	selectedYAxe int
}

func NewGame(board Board) *Game {
	return &Game{
		board:     board,
		selectedXAxe: -1,
		selectedYAxe: -1,
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()

		// Вычисляем смещение для центра поля
		offsetX, offsetY := CenterOffSet(ScreenWidth, ScreenHeight)

		tileXAxe := (cursorX - offsetX) / (tileSize + tileMargin)
		tileYAxe := (cursorY - offsetY) / (tileSize + tileMargin)

		if !(0 <= tileXAxe && tileXAxe < len(g.board[0]) && 0 <= tileYAxe && tileYAxe < len(g.board)) {
			return nil
		}

		if g.selectedXAxe == tileXAxe && g.selectedYAxe == tileYAxe {
			return nil
		}

		if g.board[tileYAxe][tileXAxe].ID == 0 {
			return nil
		}

		fmt.Println("selected:", g.selectedXAxe, g.selectedYAxe, "new:", tileXAxe, tileYAxe)
		if g.board.CanConnect(g.selectedXAxe, g.selectedYAxe, tileXAxe, tileYAxe) {
			g.board[tileYAxe][tileXAxe].Removed = true
			g.board[g.selectedYAxe][g.selectedXAxe].Removed = true
			g.selectedXAxe, g.selectedYAxe = -1, -1 
			return nil
		}

		PrintRemovedMatrix(g.board)
		g.selectedXAxe = tileXAxe
		g.selectedYAxe = tileYAxe

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	offsetX, offsetY := CenterOffSet(ScreenWidth, ScreenHeight)
	screen.Fill(color.RGBA{50, 50, 50, 255})

	for _, row := range g.board {
		for _, tile := range row {
			tile.DrawTile(screen, offsetX, offsetY, g.selectedXAxe, g.selectedYAxe)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
