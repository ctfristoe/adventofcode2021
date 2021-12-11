package problems

/// empty struct used for set membership
type void struct{}

/// typical
type Point struct{ x, y int }

func (p *Point) NeighborsInclDiagonal(xMin, xMax, yMin, yMax int) []Point {
	return append(
		p.Neighbors(xMin, xMax, yMin, yMax),
		p.DiagonalNeighbors(xMin, xMax, yMin, yMax)...,
	)
}

func (p *Point) Neighbors(xMin, xMax, yMin, yMax int) (n []Point) {
	if p.x > xMin {
		n = append(n, Point{x: p.x - 1, y: p.y})
	}
	if p.x < xMax {
		n = append(n, Point{x: p.x + 1, y: p.y})
	}
	if p.y > yMin {
		n = append(n, Point{x: p.x, y: p.y - 1})
	}
	if p.y < yMax {
		n = append(n, Point{x: p.x, y: p.y + 1})
	}
	return
}

func (p *Point) DiagonalNeighbors(xMin, xMax, yMin, yMax int) (n []Point) {
	if p.x > xMin {
		if p.y > yMin {
			n = append(n, Point{x: p.x - 1, y: p.y - 1})
		}
		if p.y < yMax {
			n = append(n, Point{x: p.x - 1, y: p.y + 1})
		}
	}
	if p.x < xMax {
		if p.y > yMin {
			n = append(n, Point{x: p.x + 1, y: p.y - 1})
		}
		if p.y < yMax {
			n = append(n, Point{x: p.x + 1, y: p.y + 1})
		}
	}
	return
}

/// RuneSet structure, pretty standard, void should use no memory
/// methods: add, contains, first, difference, intersection
type RuneSet map[rune]void

func MakeRuneSet(str string) RuneSet {
	new := make(RuneSet)
	for _, char := range str {
		new.add(char)
	}
	return new
}

func (s RuneSet) add(item rune) {
	s[item] = void{}
}

func (s RuneSet) contains(item rune) bool {
	_, found := s[item]
	return found
}

func (s RuneSet) first() rune {
	for item := range s {
		return item
	}
	panic("empty RuneSet")
}

func (s RuneSet) difference(other RuneSet) RuneSet {
	new := make(RuneSet)
	for item := range s {
		if !other.contains(item) {
			new.add(item)
		}
	}
	return new
}

func (s RuneSet) intersection(other RuneSet) RuneSet {
	new := make(RuneSet)
	for item := range s {
		if other.contains(item) {
			new.add(item)
		}
	}
	return new
}

func (s RuneSet) issubRuneSet(other RuneSet) bool {
	for item := range s {
		if !other.contains(item) {
			return false
		}
	}
	return true
}

func (s RuneSet) toString() string {
	var runes []rune
	for item := range s {
		runes = append(runes, item)
	}
	alphabetizeRunes(runes)
	return string(runes)
}
