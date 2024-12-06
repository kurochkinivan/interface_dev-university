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

type TurnsPoint struct {
	tile      *Tile
	direction Point
	turns     int
}

func NewTurnsPoint(tile *Tile, turns int) TurnsPoint {
	return TurnsPoint{
		tile:      tile,
		direction: Point{0, 0},
		turns:     turns,
	}
}

func (b Board) CanConnect(x1, y1, x2, y2 int) bool {
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return false
	}

	tile1, tile2 := b[y1][x1], b[y2][x2]
	fmt.Println("tiles", tile1, tile2)
	point1, point2 := NewTurnsPoint(tile1, 0), NewTurnsPoint(tile2, 0)

	found := BFS(b, point1, point2, 3)
	if found {
		fmt.Println("FOUND!")
	} else {
		fmt.Println("NOT FOUND:(")
	}
	return BFS(b, point1, point2, 3)
}

func BFS(graph [][]*Tile, start, end TurnsPoint, maxTurns int) bool {
	n := len(graph)
	m := len(graph[0])

	queue := []TurnsPoint{start}
	visited := make([][]bool, n)
	for row := range visited {
		visited[row] = make([]bool, m)
	}
	visited[start.tile.Position.Y][start.tile.Position.X] = true

	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.turns == maxTurns {
			continue
		}

		if end.tile.Position.X == cur.tile.Position.X && end.tile.Position.Y == cur.tile.Position.Y {
			fmt.Println("ID", start.tile.ID, end.tile.ID)
			return true
		}

		for _, dir := range directions {
			nextX := cur.tile.Position.X + dir.X
			nextY := cur.tile.Position.Y + dir.Y

			if isValid(nextX, nextY, n, m, graph, visited, end.tile) {
				newTurns := cur.turns
				if cur.direction != dir {
					newTurns++
				}
				queue = append(queue, NewTurnsPoint(graph[nextY][nextX], newTurns))
				fmt.Println(queue)
				visited[nextY][nextX] = true
			}
		}
	}

	return false
}

func isValid(x, y, n, m int, graph [][]*Tile, visited [][]bool, want *Tile) bool {
	return 0 <= x && x < n && 0 <= y && y < m && !visited[y][x] &&
		(graph[x][y].Removed || (graph[y][x].Position.X == want.Position.X && graph[y][x].Position.Y == want.Position.Y))
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
			board[y][x] = NewTile(ids[index], x, y, false)
			index++
		}
	}

	return board
}
