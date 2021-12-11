package problems

type DumboOctopus struct{}

func (DumboOctopus) DoProblemOne(filepath string) int {
	const turns = 100
	matrix := ReadIntegerMatrix(filepath)
	grid := fromIntMatrix(matrix)
	for i := 0; i < turns; i++ {
		grid.incrementAllPoints()
		grid.flashAllHighEnergyPoints()
		grid.resetAllFlashingPoints()
		grid.addCountToTotal()
	}
	return grid.total
}

func (DumboOctopus) DoProblemTwo(filepath string) int {
	matrix := ReadIntegerMatrix(filepath)
	grid := fromIntMatrix(matrix)
	// 1000 seems like when we should stop and panic
	for i := 0; i < 1000; i++ {
		grid.incrementAllPoints()
		grid.flashAllHighEnergyPoints()
		grid.resetAllFlashingPoints()
		if grid.count == 100 {
			return i + 1 // index 0 = turn 1
		}
		grid.addCountToTotal()
	}
	panic("too many turns have passed")
}

func fromIntMatrix(matrix [][]int) *energygrid {
	var energies [10][10]int
	for y, row := range matrix {
		for x, energy := range row {
			energies[y][x] = energy
		}
	}
	return &energygrid{
		energies: energies,
		flashing: make(map[Point]bool),
		count:    0,
	}
}

type energygrid struct {
	energies [10][10]int
	flashing map[Point]bool
	count    int
	total    int
}

func (*energygrid) neighbors(p Point) []Point {
	return p.NeighborsInclDiagonal(0, 9, 0, 9)
}

func (grid *energygrid) energy(p Point) int {
	return grid.energies[p.y][p.x]
}

func (grid *energygrid) increment(p Point) {
	grid.energies[p.y][p.x]++
}

func (grid *energygrid) reset(p Point) {
	grid.energies[p.y][p.x] = 0
}

func (grid *energygrid) flash(p Point) {
	if grid.flashing[p] {
		return // a flashing point can't flash again
	}
	grid.count++
	grid.flashing[p] = true
	for _, neighbor := range grid.neighbors(p) {
		grid.increment(neighbor)
		if grid.energy(neighbor) > 9 {
			grid.flash(neighbor)
		}
	}
}

func (grid *energygrid) addCountToTotal() {
	grid.total += grid.count
	grid.count = 0
}

func (grid *energygrid) incrementAllPoints() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			grid.increment(Point{x, y})
		}
	}
}

func (grid *energygrid) flashAllHighEnergyPoints() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			p := Point{x, y}
			if grid.energy(p) > 9 {
				grid.flash(p)
			}
		}
	}
}

func (grid *energygrid) resetAllFlashingPoints() {
	for p, _ := range grid.flashing {
		grid.reset(p)
	}
	grid.flashing = make(map[Point]bool)
}
