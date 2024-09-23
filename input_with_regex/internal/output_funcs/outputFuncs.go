package outputfuncs

import (
	"fmt"

	"github.com/kurochkinivan/input_with_regex/internal/parser"
)

func Print(ps []parser.Point, ls []parser.Line, cs []parser.Circle) {
	fmt.Println("----------- points -----------")
	for i := 0; i < len(ps); i++ {
		fmt.Printf("%2.2d:\tx: %.3f, y: %.3f;\n", i+1, ps[i].X, ps[i].Y)
	}

	fmt.Println("----------- lines -----------")
	for i := 0; i < len(ls); i++ {
		fmt.Printf("%2.2d:\tx1: %.3f, y1: %.3f;\tx2: %.3f, y2:%.3f\n", i+1, ls[i].Start.X, ls[i].Start.Y, ls[i].End.X, ls[i].End.Y)
	}

	fmt.Println("----------- circles -----------")
	for i := 0; i < len(cs); i++ {
		fmt.Printf("%2.2d:\tx: %.3f, y:%.3f;\tr: %.3f\n", i+1, cs[i].Point.X, cs[i].Point.Y, cs[i].Radius)
	}
}

func Count(ps []parser.Point, ls []parser.Line, cs []parser.Circle) {
	total := len(ps) + len(ls) + len(cs)
	fmt.Printf("Total count: %d\n", total)
}
