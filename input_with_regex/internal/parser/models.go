package parser

type Point struct {
	X float64
	Y float64
}

type Line struct {
	Start Point
	End   Point
}

type Circle struct {
	Point  Point
	Radius float64
}