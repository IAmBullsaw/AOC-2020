package grid

type Grid struct {
	minX, minY, maxX, maxY int
	grid                   map[Coordinate]interface{}
}

func NewGrid(minx, miny, maxx, maxy int) Grid {
	return Grid{minX: minx, minY: miny, maxX: maxx, maxY: maxy, grid: make(map[Coordinate]interface{})}
}

func (g *Grid) At(c Coordinate) interface{} {
	return g.grid[c]
}

func (g *Grid) Add(c Coordinate, val interface{}) bool {
	if g.grid[c] == nil {
		g.grid[c] = val
		return true
	}
	return false
}

func (g *Grid) Set(c Coordinate, val interface{}) {
	g.grid[c] = val
}

func (g Grid) Clone() Grid {
	grid := NewGrid(g.minX, g.minY, g.maxX, g.maxY)
	for k, v := range g.grid {
		grid.Add(k, v)
	}
	return grid
}
