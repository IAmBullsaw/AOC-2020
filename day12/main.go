package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	g "github.com/IAmBullsaw/AOC-2020/pkg/grid"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	face := 270
	start := g.NewCoordinate(0, 0)
	pos := g.NewCoordinate(0, 0)

	for _, d := range strings.Fields(puzzle) {
		action := d[0]
		val, _ := strconv.Atoi(d[1:])
		fmt.Println(pos)

		if action == 'F' {
			fmt.Printf("F -> ")
			if deg := face % 360; deg == 0 {
				action = 'E'
			} else if deg := face % 360; deg == 90 || deg == -90 {
				action = 'N'
			} else if deg := face % 360; deg == 180 || deg == -180 {
				action = 'W'
			} else if deg := face % 360; deg == 270 || deg == -270 {
				action = 'S'
			}
			fmt.Println(string(action))
		}
		fmt.Println("Action", string(action), val)
		if action == 'N' {
			pos = pos.Up(val)
		} else if action == 'S' {
			pos = pos.Down(val)
		} else if action == 'E' {
			pos = pos.Right(val)
		} else if action == 'W' {
			pos = pos.Left(val)
		} else if action == 'L' {
			face += val
		} else if action == 'R' {
			face -= val
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
		fmt.Println("WP Before:", wp)
		action := d[0]
		val, _ := strconv.Atoi(d[1:])
		if action == 'F' {
			diff := wp.Subtract(pos)
			fmt.Println(" F Before:", pos, wp, diff)

			pos = g.NewCoordinate(diff.X()*val+pos.X(), diff.Y()*val+pos.Y())
			wp = pos.Add(diff)
			fmt.Println(" F After: ", pos, wp, diff)
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
			if wp.X() > 0 && wp.Y() > 0 || wp.X() < 0 && wp.Y() < 0 {
				wp = g.NewCoordinate(wp.X()*-1, wp.Y())
			} else { //if wp.X() > 0 && wp.Y() < 0 || wp.X() < 0 && wp.Y() > 0 {
				wp = g.NewCoordinate(wp.X(), wp.Y()*-1)
			}
		} else if action == 'R' {
			if wp.X() > 0 && wp.Y() > 0 || wp.X() < 0 && wp.Y() < 0 {
				wp = g.NewCoordinate(wp.X(), wp.Y()*-1)
			} else { //if wp.X() > 0 && wp.Y() < 0 || wp.X() < 0 && wp.Y() > 0 {
				wp = g.NewCoordinate(wp.X()*-1, wp.Y())
			}
		}
		fmt.Println("WP After:  ", wp)
	}
	return start.ManhattanDistance(pos)
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
