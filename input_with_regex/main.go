package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var str string = `(^Point\(([-+]?\d+\.?\d*),\s([-+]?\d+\.?\d*)\))|(Line\(Point\(([-+]?\d+\.?\d*),\s([-+]?\d+\.?\d*)\), Point\(([-+]?\d+\.?\d*),\s([-+]?\d+\.?\d*)\)\))|(Circle\(Point\(([-+]?\d+\.?\d*),\s([-+]?\d+\.?\d*)\), (\+?\d+\.?\d*)\))`

var (
	float       = `([-+]?\d+\.?\d*)`
	posFloat    = `(\+?\d+\.?\d*)`
	point       = fmt.Sprintf(`Point\(%s,\s%[1]s\)`, float)
	pointRegex  = fmt.Sprintf(`^%s`, point)
	lineRegex   = fmt.Sprintf(`Line\(%s, %[1]s\)`, point)
	circleRegex = fmt.Sprintf(`Circle\(%s, %s\)`, point, posFloat)
	regex       = fmt.Sprintf(`(%s)|(%s)|(%s)`, pointRegex, lineRegex, circleRegex)
)

func main() {
	fileName := "./data.txt"

	print := func(ps []Point, ls []Line, cs []Circle) {
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

	count := func(ps []Point, ls []Line, cs []Circle) {
		total := len(ps) + len(ls) + len(cs)
		fmt.Printf("Total count: %d\n", total)
	}

	err := ParseObjectsFromFile(fileName, print)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = ParseObjectsFromFile(fileName, count)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func ParseObjectsFromFile(fileName string, f func([]Point, []Line, []Circle)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file, err: %v", err)
	}

	pointRgx := regexp.MustCompile(pointRegex)
	lineRgx := regexp.MustCompile(lineRegex)
	circleRgx := regexp.MustCompile(circleRegex)

	var points []Point
	var lines []Line
	var circles []Circle

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		if pointRgx.MatchString(line) {
			match := pointRgx.FindStringSubmatch(line)

			vals, err := convertValuesInFloat(match)
			if err != nil {
				return err
			}
			x, y := vals[0], vals[1]

			p := Point{
				X: x,
				Y: y,
			}

			points = append(points, p)
			continue
		}

		if lineRgx.MatchString(line) {
			match := lineRgx.FindStringSubmatch(line)
			vals, err := convertValuesInFloat(match)
			if err != nil {
				return err
			}
			x1, y1, x2, y2 := vals[0], vals[1], vals[2], vals[3]

			l := Line{
				Start: Point{
					X: x1,
					Y: y1,
				},
				End: Point{
					X: x2,
					Y: y2,
				},
			}
			lines = append(lines, l)
			continue
		}

		if circleRgx.MatchString(line) {
			match := circleRgx.FindStringSubmatch(line)

			vals, err := convertValuesInFloat(match)
			if err != nil {
				return err
			}
			x, y, r := vals[0], vals[1], vals[2]

			c := Circle{
				Point: Point{
					X: x,
					Y: y,
				},
				Radius: r,
			}
			circles = append(circles, c)
			continue
		}
	}

	f(points, lines, circles)

	return nil
}

func convertValuesInFloat(match []string) ([]float64, error) {
	values := make([]float64, len(match))
	for i := 1; i < len(match); i++ {
		v, err := strconv.ParseFloat(match[i], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float, str: %s, err: %v", match[i], err)
		}
		values[i-1] = v
	}

	return values, nil
}

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
