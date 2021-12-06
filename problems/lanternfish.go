package problems

type Lanternfish struct{}

func (Lanternfish) DoProblemOne(filepath string) int {
	const N = 80
	input := ReadCommaSeparatedIntegers(filepath)
	return getFishAfterNGenerations(input, N)
}

func (Lanternfish) DoProblemTwo(filepath string) int {
	const N = 256
	input := ReadCommaSeparatedIntegers(filepath)
	return getFishAfterNGenerationsPerformant(input, N)
}

type lanternfish struct {
	timer int
}

func (l *lanternfish) tick() []lanternfish {
	if l.timer == 0 {
		return []lanternfish{{6}, {8}}
	}
	return []lanternfish{{l.timer - 1}}
}

func getFishAfterNGenerationsPerformant(initial []int, n int) int {
	timers := getTimerMapFromPuzzleInput(initial)
	for i := 0; i < n; i++ {
		timers = getNextGenerationMap(timers)
	}
	return countFishInMap(timers)
}

func countFishInMap(m map[int]int) int {
	total := 0
	for _, count := range m {
		total += count
	}
	return total
}

func getFishAfterNGenerations(initial []int, n int) int {
	fish := getLanternfishFromPuzzleInput(initial)
	for i := 0; i < n; i++ {
		fish = getNextGeneration(fish)
	}
	return len(fish)
}

func getNextGenerationMap(m map[int]int) map[int]int {
	new := make(map[int]int, 9)
	reproductive := m[0]
	new[6] += reproductive
	new[8] += reproductive
	for i := 1; i < 9; i++ {
		new[i-1] += m[i]
	}
	return new
}

func getNextGeneration(fishes []lanternfish) (new []lanternfish) {
	for _, fish := range fishes {
		new = append(new, fish.tick()...)
	}
	return
}

func getTimerMapFromPuzzleInput(input []int) map[int]int {
	m := make(map[int]int, 9)
	for _, timer := range input {
		m[timer]++
	}
	return m
}

func getLanternfishFromPuzzleInput(input []int) (fish []lanternfish) {
	for _, timer := range input {
		fish = append(fish, lanternfish{timer})
	}
	return
}
