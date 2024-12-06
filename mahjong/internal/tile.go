package internal

import (
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	tileSize   = 60
	tileMargin = 8
	borderSize = 4
)

type Tile struct {
	ID       int
	Position Point
	Removed  bool
	Image    *ebiten.Image
}

type Point struct {
	X int
	Y int
}

func NewTile(id int, x, y int, removed bool) *Tile {
	return &Tile{
		ID: id,
		Position: Point{
			X: x,
			Y: y,
		},
		Removed: removed,
		Image:   generateImage(id),
	}
}

func generateImage(id int) *ebiten.Image {
	img := ebiten.NewImage(tileSize, tileSize)
	img.Fill(color.White)

	face := basicfont.Face7x13
	idStr := strconv.Itoa(id)

	textBounds, _ := font.BoundString(face, idStr)
	textWidth := (textBounds.Max.X - textBounds.Min.X).Ceil()
	textHeight := (textBounds.Max.Y - textBounds.Min.Y).Ceil()

	x := (tileSize - textWidth) / 2
	y := (tileSize-textHeight)/2 + textHeight

	text.Draw(img, idStr, face, x, y, color.Black)

	return img
}

func (t *Tile) DrawTile(screen *ebiten.Image, offsetX, offsetY int, selectedX, selectedY int) {
	if t.Removed {
		return
	}

	x := float64(offsetX) + float64(t.Position.X)*(tileSize+tileMargin)
	y := float64(offsetY) + float64(t.Position.Y)*(tileSize+tileMargin)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(t.Image, op)

	if int(t.Position.X) == selectedX && int(t.Position.Y) == selectedY {
		border := ebiten.NewImage(tileSize+borderSize, tileSize+borderSize)

		// Верхняя и нижняя границы
		border.SubImage(image.Rect(0, 0, tileSize+borderSize, borderSize)).(*ebiten.Image).Fill(red)
		border.SubImage(image.Rect(0, tileSize, tileSize+borderSize, tileSize+borderSize)).(*ebiten.Image).Fill(red)

		// Правая и левая границы
		border.SubImage(image.Rect(0, 0, borderSize, tileSize+borderSize)).(*ebiten.Image).Fill(red)
		border.SubImage(image.Rect(tileSize, 0, tileSize+borderSize, tileSize+borderSize)).(*ebiten.Image).Fill(red)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x-borderSize/2, y-borderSize/2)
		screen.DrawImage(border, op)
	}
}
