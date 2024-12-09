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
	Tile  *Tile
	Dir   Direction
	Turns int
}

func NewTurnsPoint(tile *Tile, dir Direction, turns int) TurnsPoint {
	return TurnsPoint{
		Tile:  tile,
		Dir:   dir,
		Turns: turns,
	}
}

func (b Board) CanConnect(x1, y1, x2, y2 int) bool {
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return false
	}

	tile1, tile2 := b[y1][x1], b[y2][x2]
	fmt.Println("tiles", tile1, tile2)
	point1, point2 := NewTurnsPoint(tile1, Direction{0, 0}, 0), NewTurnsPoint(tile2, Direction{0, 0}, 0)

	found := BFS(b, point1, point2, 3)
	if found {
		fmt.Println("FOUND!")
	} else {
		fmt.Println("NOT FOUND:(")
	}
	return found
}

// TODO: think of borders of graph
func BFS(graph [][]*Tile, start, end TurnsPoint, maxTurns int) bool {
	n := len(graph)
	m := len(graph[0])

	queue := []TurnsPoint{NewTurnsPoint(start.Tile, Direction{0, 0}, 0)}
	visited := make([][]bool, n)
	for row := range visited {
		visited[row] = make([]bool, m)
	}
	visited[start.Tile.Position.Y][start.Tile.Position.X] = true

	directions := []Direction{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if end.Tile.Position.X == cur.Tile.Position.X && end.Tile.Position.Y == cur.Tile.Position.Y {
			fmt.Println("ID", start.Tile.ID, end.Tile.ID)
			return start.Tile.ID == end.Tile.ID
		}

		if cur.Turns == maxTurns {
			continue
		}

		for _, dir := range directions {
			nextX := cur.Tile.Position.X + dir.X
			nextY := cur.Tile.Position.Y + dir.Y

			if isValid(nextX, nextY, n, m, graph, visited, end.Tile) {
				nextTurns := cur.Turns
				if cur.Dir != dir {
					nextTurns++
				}

				queue = append(queue, NewTurnsPoint(graph[nextY][nextX], dir, nextTurns))
				visited[nextY][nextX] = true
				// fmt.Println(dir)
				// fmt.Println(queue)
			}
		}
	}

	// fmt.Println("not found, end of bfs")

	return false
}

func isValid(x, y, n, m int, graph [][]*Tile, visited [][]bool, want *Tile) bool {
	return 0 <= x && x < m && 0 <= y && y < n && !visited[y][x] &&
		(graph[y][x].Removed || (graph[y][x].Position.X == want.Position.X && graph[y][x].Position.Y == want.Position.Y))
}

func InitializeBoard() Board {
	board := make(Board, rows+1)
	for i := range board {
		board[i] = make([]*Tile, cols+1)
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
	for y := 1; y < rows; y++ {
		for x := 1; x < cols; x++ {
			board[y][x] = NewTile(ids[index], x, y, false)
			index++
		}
		board[y][0] = NewTile(0, 0, y, true)
		board[y][cols] = NewTile(0, cols, y, true)
	}

	for col := range board[0] {
		board[0][col] = NewTile(0, col, 0, true)
		board[rows][col] = NewTile(0, col, 0, true)
	}

	PrintRemovedMatrix(board)
	return board
}

func PrintRemovedMatrix(board Board) {
	for _, row := range board {
		for _, tile := range row {
			if tile != nil {
				// fmt.Printf("%.2d ", tile.ID)
				if tile.Removed {
					fmt.Print("0 ")
				} else {
					fmt.Print("1 ")
				}
			} else {
				// Если плитка отсутствует, можно вывести символ для обозначения пустоты
				fmt.Print("x ")
			}
		}
		fmt.Println() // Новый ряд
	}
}
