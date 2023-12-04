package main

import (
	"github.com/alc6/aoc2023/util"
	"log"
	"math"
	"regexp"
	"strconv"
)

func main() {
	lines, err := util.ReadFileLines("d4/input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	scratchedCards := make([]*ScratchedCard, 0, len(lines))
	for _, line := range lines {
		scratchedCards = append(scratchedCards, NewScratchedCard(line))
	}

	pt1(scratchedCards)
	pt2(scratchedCards)
}

type ScratchedCard struct {
	ID             int
	PlayedNumbers  map[int]struct{}
	WinningNumbers map[int]struct{}
}

func NewScratchedCard(in string) *ScratchedCard {
	var (
		cr = ScratchedCard{
			ID:             0,
			PlayedNumbers:  make(map[int]struct{}),
			WinningNumbers: make(map[int]struct{}),
		}
		re = regexp.MustCompile(`Card\s+(\d+):\s*([^|]+)\|\s*(.+)`)
	)

	m := re.FindStringSubmatch(in)

	id, _ := strconv.Atoi(m[1])
	cr.ID = id

	playedNumbers := getNumbers(m[2])
	for _, n := range playedNumbers {
		cr.PlayedNumbers[n] = struct{}{}
	}

	winningNumbers := getNumbers(m[3])
	for _, n := range winningNumbers {
		cr.WinningNumbers[n] = struct{}{}
	}

	return &cr
}

func (cr ScratchedCard) TotalPoints() int {
	var matches int
	for n := range cr.PlayedNumbers {
		if _, ok := cr.WinningNumbers[n]; ok {
			matches++
		}
	}

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func (cr ScratchedCard) MatchingNumbers() int {
	matchingNumbers := make([]string, 0, len(cr.PlayedNumbers))
	for n := range cr.PlayedNumbers {
		if _, ok := cr.WinningNumbers[n]; ok {
			matchingNumbers = append(matchingNumbers, strconv.Itoa(n))
		}
	}

	return len(matchingNumbers)
}

func pt1(scratchedCards []*ScratchedCard) {
	var totalPoint int
	for _, sc := range scratchedCards {
		totalPoint += sc.TotalPoints()
	}

	log.Println("pt1 total points: ", totalPoint)
}

func pt2(scratchedCards []*ScratchedCard) {
	cardCopies := make([]int, len(scratchedCards))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	for i, sc := range scratchedCards {
		matchingNumbers := sc.MatchingNumbers()

		for j := i + 1; j < len(scratchedCards) && j <= i+matchingNumbers; j++ {
			cardCopies[j] += cardCopies[i]
		}
	}

	var totalCards int
	for _, copies := range cardCopies {
		totalCards += copies
	}

	log.Println("pt2 total scratchcards: ", totalCards)
}

func getNumbers(l string) []int {
	var (
		numbers = make([]int, 0)
		re      = regexp.MustCompile(`\d+`)
	)

	for _, n := range re.FindAllString(l, -1) {
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	return numbers
}
