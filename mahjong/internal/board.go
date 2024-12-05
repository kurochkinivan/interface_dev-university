package internal

import (
	"fmt"
	"math/rand/v2"
)

const (
	rows = 10
	cols = 14
)

const (
	fieldWidth  = (cols * tileSize) + ((cols - 1) * tileMargin)
	fieldHeight = (rows * tileSize) + ((rows - 1) * tileMargin)
)

func CenterOffSet(screenWidth, screenHeight int) (offsetX, offsetY int) {
	offsetX = (screenWidth - fieldWidth) / 2
	offsetY = (screenHeight - fieldHeight) / 2
	return offsetX, offsetY
}

type Board [][]*Tile

func (b Board) CanConnect(tile1X, tile1Y, tile2X, tile2Y int) bool {
	if tile1X == tile2X || tile1X < 0 || tile1Y < 0 || tile2X < 0 || tile2Y < 0 {
		return false
	}

	tile1 := b[tile1Y][tile1X]
	tile2 := b[tile2Y][tile2X]
	fmt.Println(tile1, tile2)

	return false
}

type PathState struct {
	X, Y    int
	Turns   int
	Path    []TilePosition
	LastDir int
}

func InitializeBoard() Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]*Tile, cols)
	}

	ids := make([]int, 0, cols*rows)
	for id := 1; id <= 14; id++ {
		for i := 0; i < 10; i++ {
			ids = append(ids, id)
		}
	}

	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	index := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			board[y][x] = NewTile(ids[index], float64(x), float64(y), false)
			index++
		}
	}

	return board
}
