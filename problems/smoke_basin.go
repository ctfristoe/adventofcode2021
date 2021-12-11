package problems

import "sort"

type SmokeBasin struct{}

func (SmokeBasin) DoProblemOne(filepath string) (riskLevel int) {
	input := ReadLines(filepath)
	heightmap := NewHeightmap(input)
	for _, p := range heightmap.getLocalMinima() {
		riskLevel += heightmap.height(p) + 1
	}
	return
}

func (SmokeBasin) DoProblemTwo(filepath string) (product int) {
	input := ReadLines(filepath)
	heightmap := NewHeightmap(input)
	var sizes []int
	for _, p := range heightmap.getLocalMinima() {
		sizes = append(sizes, heightmap.getBasinSize(p))
	}
	sort.Ints(sizes)
	return sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]
}

type heightmap [][]int

func (hm heightmap) width() int {
	return len(hm[0])
}

func (hm heightmap) length() int {
	return len(hm)
}

func (hm heightmap) height(p Point) int {
	return hm[p.y][p.x]
}

func (hm heightmap) neighbors(p Point) (neighbors []Point) {
	if p.x > 0 {
		neighbors = append(neighbors, Point{x: p.x - 1, y: p.y})
	}
	if p.x+1 < hm.width() {
		neighbors = append(neighbors, Point{x: p.x + 1, y: p.y})
	}
	if p.y > 0 {
		neighbors = append(neighbors, Point{x: p.x, y: p.y - 1})
	}
	if p.y+1 < hm.length() {
		neighbors = append(neighbors, Point{x: p.x, y: p.y + 1})
	}
	return
}

func (hm heightmap) isLocalMin(p Point) bool {
	thisHeight := hm.height(p)
	for _, other := range hm.neighbors(p) {
		if thisHeight >= hm.height(other) {
			return false
		}
	}
	return true
}

func (hm heightmap) getBasinSize(p Point) (size int) {
	basin := make(map[Point]bool)
	hm.addNearbyBasinPoints(p, basin)
	return len(basin)
}

func (hm heightmap) addNearbyBasinPoints(p Point, basin map[Point]bool) {
	basin[p] = true
	for _, other := range hm.neighbors(p) {
		if basin[other] {
			continue
		}
		if hm.height(other) == 9 {
			continue
		}
		hm.addNearbyBasinPoints(other, basin)
	}
}

func (hm heightmap) getLocalMinima() (Points []Point) {
	for x := 0; x < hm.width(); x++ {
		for y := 0; y < hm.length(); y++ {
			p := Point{x, y}
			if hm.isLocalMin(p) {
				Points = append(Points, p)
			}
		}
	}
	return
}

func NewHeightmap(input []string) heightmap {
	hm := make(heightmap, len(input))
	for i, str := range input {
		hm[i] = stringToIntegerSlice(str)
	}
	return hm
}

func stringToIntegerSlice(str string) (slice []int) {
	for _, char := range str {
		slice = append(slice, int(char-'0'))
	}
	return
}
