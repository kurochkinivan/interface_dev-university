package internal

type Direction Point 

func (d Direction) String() string {
	switch d {
	case Direction(Point{-1, 0}):
		return "вверх"
	case Direction(Point{1, 0}):
		return "вниз"
	case Direction(Point{0, 1}):
		return "вправо"
	case Direction(Point{0, -1}):
		return "влево"
	case Direction(Point{0, 0}):
		return "пустое направление"
	default:
		panic("unknown direction")
	}
}