package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/kurochkinivan/input_with_regex/internal/constants"
)


func ParseObjectsFromFile(fileName string, f func([]Point, []Line, []Circle)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file, err: %v", err)
	}
	
	pointRgx := regexp.MustCompile(constants.PointRegex)
	lineRgx := regexp.MustCompile(constants.LineRegex)
	circleRgx := regexp.MustCompile(constants.CircleRegex)

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
