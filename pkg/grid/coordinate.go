package grid

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d:%d)", c.x, c.y)
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

func (c Coordinate) X() int {
	return c.x
}

func (c Coordinate) Y() int {
	return c.y
}

func (c Coordinate) Up(v int) Coordinate {
	return Coordinate{x: c.x, y: c.y - v}
}

func (c Coordinate) Down(v int) Coordinate {
	return Coordinate{x: c.x, y: c.y + v}
}

func (c Coordinate) Left(v int) Coordinate {
	return Coordinate{x: c.x - v, y: c.y}
}
func (c Coordinate) Right(v int) Coordinate {
	return Coordinate{x: c.x + v, y: c.y}
}

func (c Coordinate) In(grid Grid) bool {
	xin := c.x >= grid.minX && c.x <= grid.maxX
	yin := c.y >= grid.minY && c.y <= grid.maxY
	return xin && yin
}

func (p1 Coordinate) Add(p2 Coordinate) Coordinate {
	return Coordinate{x: p1.x + p2.x, y: p1.y + p2.y}
}

func (p1 Coordinate) Subtract(p2 Coordinate) Coordinate {
	return Coordinate{x: p1.x - p2.x, y: p1.y - p2.y}
}

func (p1 Coordinate) ManhattanDistance(p2 Coordinate) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
