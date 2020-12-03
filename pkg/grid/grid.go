package grid

type Grid struct {
	xMax, yMax int
	grid       [][]int
}

func NewGrid(grid [][]int) Grid {
	return Grid{xMax: len(grid[0]), yMax: len(grid), grid: grid}
}

func (g Grid) At(x, y int) int {
	return g.grid[y][x]
}
