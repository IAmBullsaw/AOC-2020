package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	g "github.com/IAmBullsaw/AOC-2020/pkg/grid"
)

const (
	North = 0
	East  = 90
	South = 180
	West  = 270
)

func dirStr(d int) string {
	n := map[int]string{0: "N", 90: "E", 180: "S", 270: "W"}
	return n[int(math.Abs(float64(d)))]
}

func dir(f int) int {
	if f < 0 {
		return 360 + f
	}
	return f
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	face := East
	start := g.NewCoordinate(0, 0)
	pos := g.NewCoordinate(0, 0)

	for _, d := range strings.Fields(puzzle) {
		action := d[0]
		val, _ := strconv.Atoi(d[1:])

		if action == 'F' {
			switch dir(face) {
			case North:
				action = 'N'
				break
			case East:
				action = 'E'
				break
			case South:
				action = 'S'
				break
			case West:
				action = 'W'
				break
			}
		}
		if action == 'N' {
			pos = pos.Up(val)
		} else if action == 'S' {
			pos = pos.Down(val)
		} else if action == 'E' {
			pos = pos.Right(val)
		} else if action == 'W' {
			pos = pos.Left(val)
		} else if action == 'L' {
			face = (face - val) % 360
		} else if action == 'R' {
			face = (face + val) % 360
		}
	}

	return start.ManhattanDistance(pos)
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	start := g.NewCoordinate(0, 0)
	pos := g.NewCoordinate(0, 0)
	wp := g.NewCoordinate(10, 1)

	for _, d := range strings.Fields(puzzle) {
		action := d[0]
		val, _ := strconv.Atoi(d[1:])

		if action == 'F' {
			pos.X = pos.X + (wp.X * val)
			pos.Y = pos.Y + (wp.Y * val)
		}
		if action == 'N' {
			wp = wp.Up(val)
		} else if action == 'S' {
			wp = wp.Down(val)
		} else if action == 'E' {
			wp = wp.Right(val)
		} else if action == 'W' {
			wp = wp.Left(val)
		} else if action == 'L' {
			turns := val / 90
			for turns > 0 {
				wp.X, wp.Y = wp.Y*-1, wp.X
				turns--
			}
		} else if action == 'R' {
			turns := val / 90
			for turns > 0 {
				wp.X, wp.Y = wp.Y, wp.X*-1
				turns--
			}
		}
	}

	return start.ManhattanDistance(pos)
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
