package problems

type Whales struct{}

func (Whales) DoProblemOne(filepath string) int {
	input := ReadCommaSeparatedIntegers(filepath)
	estimator := newFuelEstimator(input)
	return estimator.getMinimalRequiredFuel()
}

func (Whales) DoProblemTwo(filepath string) int {
	input := ReadCommaSeparatedIntegers(filepath)
	estimator := newFuelEstimator(input)
	return estimator.getMinimalRequiredFuelComplex()
}

type estimator struct {
	table map[int]int
	min   int
	max   int
}

func (e *estimator) getMinimalRequiredFuelComplex() int {
	minFuel := e.estimateFuelComplex(e.min)
	for i := e.min + 1; i <= e.max; i++ {
		fuel := e.estimateFuelComplex(i)
		minFuel = Min(fuel, minFuel)
	}
	return minFuel
}

func (e *estimator) getMinimalRequiredFuel() int {
	minFuel := e.estimateFuel(e.min)
	for i := e.min + 1; i <= e.max; i++ {
		fuel := e.estimateFuel(i)
		minFuel = Min(fuel, minFuel)
	}
	return minFuel
}

func (e *estimator) estimateFuelComplex(pos int) int {
	total := 0
	for crabPos, count := range e.table {
		distance := AbsDiff(pos, crabPos)
		total += count * NthTriangleNumber(distance)
	}
	return total
}

func (e *estimator) estimateFuel(pos int) int {
	total := 0
	for key, value := range e.table {
		total += value * AbsDiff(pos, key)
	}
	return total
}

func newFuelEstimator(input []int) estimator {
	n := input[0]
	table := map[int]int{n: 1}
	min, max := n, n
	for _, pos := range input[1:] {
		min = Min(pos, min)
		max = Max(pos, max)
		table[pos]++
	}
	return estimator{table, min, max}
}
