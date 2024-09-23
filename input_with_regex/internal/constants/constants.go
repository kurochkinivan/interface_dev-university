package constants

import "fmt"

const (
	float    = `([-+]?\d+\.?\d*)`
	posFloat = `(\+?\d+\.?\d*)`
	Version  = "1.0.0"
)

var (
	point = fmt.Sprintf(`Point\(%s,\s%[1]s\)`, float)

	PointRegex  = fmt.Sprintf(`^%s`, point)
	LineRegex   = fmt.Sprintf(`Line\(%s, %[1]s\)`, point)
	CircleRegex = fmt.Sprintf(`Circle\(%s, %s\)`, point, posFloat)
)
