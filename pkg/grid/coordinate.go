package grid

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d:%d)", c.X, c.Y)
}

func NewCoordinate(X, Y int) Coordinate {
	return Coordinate{X: X, Y: Y}
}

func (c Coordinate) Up(v int) Coordinate {
	return Coordinate{X: c.X, Y: c.Y + v}
}

func (c Coordinate) Down(v int) Coordinate {
	return Coordinate{X: c.X, Y: c.Y - v}
}

func (c Coordinate) Left(v int) Coordinate {
	return Coordinate{X: c.X - v, Y: c.Y}
}
func (c Coordinate) Right(v int) Coordinate {
	return Coordinate{X: c.X + v, Y: c.Y}
}

func (c Coordinate) In(grid Grid) bool {
	xin := c.X >= grid.minX && c.X <= grid.maxX
	yin := c.Y >= grid.minY && c.Y <= grid.maxY
	return xin && yin
}

func (p1 Coordinate) Add(p2 Coordinate) Coordinate {
	return Coordinate{X: p1.X + p2.X, Y: p1.Y + p2.Y}
}

func (p1 Coordinate) Subtract(p2 Coordinate) Coordinate {
	return Coordinate{X: p1.X - p2.X, Y: p1.Y - p2.Y}
}

func (p1 Coordinate) ManhattanDistance(p2 Coordinate) int {
	return int(math.Abs(float64(p1.Y-p2.Y))) + int(math.Abs(float64(p1.X-p2.X)))
}
