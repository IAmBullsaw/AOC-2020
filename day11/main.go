package main

import (
	"fmt"
	"strings"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

const (
	Floor = iota
	Free
	Occupied
)

func adjacent(x, y int, grid [][]int) (adj int) {
	maxY := len(grid)
	maxX := len(grid[0])
	minY := 0
	minX := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			if cmp.InBounds(minX, x+dx, maxX-1, true) == cmp.Within {
				if cmp.InBounds(minY, y+dy, maxY-1, true) == cmp.Within {
					if grid[y+dy][x+dx] != Occupied {
						adj++
					}
				} else {
					adj++
				}

			} else {
				adj++
			}
		}
	}
	return
}

func print(grid [][]int) {
	for _, r := range grid {
		for _, c := range r {
			if c == Floor {
				fmt.Printf(".")
			} else if c == Free {
				fmt.Printf("L")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

func step(grid [][]int) ([][]int, bool) {
	next := [][]int{}
	changed := false
	for y, row := range grid {
		r := []int{}
		for x, c := range row {
			if c == Floor {
				r = append(r, Floor)
				continue
			}
			adj := adjacent(x, y, grid)
			// fmt.Println(x, y, "adj", adj)
			if c == Free && adj == 8 {
				r = append(r, Occupied)
				changed = true

			} else if c == Occupied && adj <= 4 {
				r = append(r, Free)
				changed = true
			} else {
				r = append(r, c)
			}
		}
		next = append(next, r)
	}
	return next, changed
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	grid := [][]int{}
	for _, row := range strings.Fields(puzzle) {
		r := []int{}
		for _, c := range row {
			if c == '.' {
				r = append(r, Floor)
			} else if c == 'L' {
				r = append(r, Free)
			} else if c == '#' {
				r = append(r, Occupied)
			}
		}
		grid = append(grid, r)
	}

	counter := 1
	for {
		next, changed := step(grid)
		if !changed {
			for _, r := range next {
				for _, s := range r {
					if s == Occupied {
						answer++
					}
				}
			}
			break
		}
		grid = next
		counter++
	}
	return

}

func adjacentSearch(px, py int, grid [][]int) (adj int) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			x := px + dx
			y := py + dy
			for cmp.InBounds(0, x, len(grid[0])-1, true) == cmp.Within && cmp.InBounds(0, y, len(grid)-1, true) == cmp.Within && grid[y][x] == Floor {
				x += dx
				y += dy
			}
			if cmp.InBounds(0, x, len(grid[0])-1, true) == cmp.Within && cmp.InBounds(0, y, len(grid)-1, true) == cmp.Within {
				// fmt.Println(grid[y][x])
				if grid[y][x] == Occupied {
					adj++
				}
			}

		}
	}
	return
}

func step2(grid [][]int) ([][]int, bool) {
	next := [][]int{}
	changed := false
	for y, row := range grid {
		r := []int{}
		for x, c := range row {
			if c == Floor {
				r = append(r, Floor)
				continue
			}
			adj := adjacentSearch(x, y, grid)
			// fmt.Println(x, y, "adj", adj)
			if c == Free && adj == 0 {
				r = append(r, Occupied)
				changed = true
			} else if c == Occupied && 5 <= adj {
				r = append(r, Free)
				changed = true
			} else {
				r = append(r, c)
			}
		}
		next = append(next, r)
	}
	return next, changed
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	grid := [][]int{}
	for _, row := range strings.Fields(puzzle) {
		r := []int{}
		for _, c := range row {
			if c == '.' {
				r = append(r, Floor)
			} else if c == 'L' {
				r = append(r, Free)
			} else if c == '#' {
				r = append(r, Occupied)
			}
		}
		grid = append(grid, r)
	}

	counter := 1
	for {
		next, changed := step2(grid)
		if !changed {
			for _, r := range next {
				for _, s := range r {
					if s == Occupied {
						answer++
					}
				}
			}
			break
		}
		grid = next
		counter++
	}
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
