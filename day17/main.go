package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

const (
	Active   = '#'
	Inactive = '.'
)

type Coordinate struct {
	X, Y, Z int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%3d,%3d,%3d)", c.X, c.Y, c.Z)
}

type PocketDimension struct {
	Zmin, Xmin, Ymin, Zmax, Xmax, Ymax int
	Active                             map[Coordinate]interface{}
	ToActivate, ToDeactivate           []Coordinate
}

func NewPocketDimension() PocketDimension {
	return PocketDimension{
		Active:       make(map[Coordinate]interface{}),
		ToActivate:   make([]Coordinate, 0),
		ToDeactivate: make([]Coordinate, 0),
		Zmin:         0, Xmin: 0, Ymin: 0, Zmax: 0, Xmax: 0, Ymax: 0}
}

func (pd *PocketDimension) Activate(c Coordinate) {
	pd.ToActivate = append(pd.ToActivate, c)
}

func (pd *PocketDimension) IsActive(c Coordinate) bool {
	_, ok := pd.Active[c]
	return ok
}

func (pd *PocketDimension) Deactivate(c Coordinate) {
	pd.ToDeactivate = append(pd.ToDeactivate, c)
}

func (pd *PocketDimension) Commit() {
	for _, c := range pd.ToDeactivate {
		delete(pd.Active, c)
	}
	fmt.Println("kept active: ", len(pd.Active))
	for _, c := range pd.ToActivate {
		pd.Active[c] = true
	}
	pd.ToDeactivate = make([]Coordinate, 0)
	pd.ToActivate = make([]Coordinate, 0)
}

func (pd *PocketDimension) UpdateBounds() {
	pd.Zmin -= 1
	pd.Zmax += 1
	pd.Xmin -= 1
	pd.Xmax += 1
	pd.Ymin -= 1
	pd.Ymax += 1
}

func (pd *PocketDimension) ActiveNeighbours(c Coordinate) (an int) {
	for dz := -1; dz < 2; dz++ {
		for dy := -1; dy < 2; dy++ {
			for dx := -1; dx < 2; dx++ {
				if dz == 0 && dy == 0 && dx == 0 {
					continue
				}
				c2 := Coordinate{X: c.X + dx, Y: c.Y + dy, Z: c.Z + dz}
				if pd.IsActive(c2) {
					an++
				}
			}
		}
	}
	return
}

func (pd *PocketDimension) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("PocketDimension (%2d,%2d,%2d):(%2d,%2d,%2d)\n", pd.Xmin, pd.Ymin, pd.Zmin, pd.Xmax, pd.Ymax, pd.Zmax))
	for z := pd.Zmin; z <= pd.Zmax; z++ {
		buf.WriteString(fmt.Sprintf("z=%d\n", z))
		for y := pd.Ymin; y <= pd.Ymax; y++ {
			for x := pd.Xmin; x <= pd.Xmax; x++ {
				coord := Coordinate{X: x, Y: y, Z: z}
				if pd.IsActive(coord) {
					// fmt.Println(coord)
					buf.WriteRune('#')
				} else {
					buf.WriteRune('.')
				}
			}
			buf.WriteRune('\n')
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}

// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.

// TODO: Loop over maxX, maxY, maxZ
func cycle(pd *PocketDimension) {
	pd.UpdateBounds()
	for z := pd.Zmin; z <= pd.Zmax; z++ {
		for y := pd.Ymin; y <= pd.Ymax; y++ {
			for x := pd.Xmin; x <= pd.Xmax; x++ {
				coord := Coordinate{X: x, Y: y, Z: z}
				an := pd.ActiveNeighbours(coord)
				if pd.IsActive(coord) && (an == 2 || an == 3) {
					// remain active
				} else {
					pd.Deactivate(coord)
				}
				if !pd.IsActive(coord) && an == 3 {
					pd.Activate(coord)
				}
			}
		}
	}
	pd.Commit()
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	pd := NewPocketDimension()
	data := strings.Fields(puzzle)
	pd.Xmax = len(data[0]) - 1
	pd.Ymax = len(data) - 1
	z := 0
	for y, line := range data {
		for x, cube := range line {
			// fmt.Println(cube)
			if cube == Active {
				pd.Activate(Coordinate{X: x, Y: y, Z: z})
			}
		}
	}
	pd.Commit()

	// fmt.Println(pd.String())
	for i := 0; i < 6; i++ {
		cycle(&pd)
		// fmt.Println(pd.String())
	}

	return len(pd.Active)
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
