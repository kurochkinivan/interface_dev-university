package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kurochkinivan/mahjong/internal"
)

func main() {
	board := internal.InitializeBoard()
	game := internal.NewGame(board)

	ebiten.SetWindowSize(internal.ScreenWidth, internal.ScreenHeight)
	ebiten.SetWindowTitle("Mahjong-Connect")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
