package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/alc6/aoc2023/util"
)

func main() {
	lines, err := util.ReadFileLines("d1/input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	digitsPerLine := make([][]int, len(lines))
	for i, l := range lines {
		digitsPerLine[i] = getDigits(l)
	}

	concatenatedDigits := make([]int, len(digitsPerLine))
	for i, digits := range digitsPerLine {
		concatenatedDigits[i] = concatDigits(digits)
	}

	var sum int
	for _, d := range concatenatedDigits {
		sum += d
	}

	log.Println("sum: ", sum)
}

func concatDigits(digits []int) int {
	if len(digits) == 0 {
		return 0
	}

	concat := 0
	for _, d := range digits {
		concat *= 10
		concat += d
	}

	return concat
}

func getDigitsAsString(s string) (int, bool) {
	var extracted strings.Builder
	// Read until a number is in the string
	for _, r := range s {
		if unicode.IsDigit(r) {
			break
		}

		extracted.WriteRune(r)

		value, exists := util.DigitsStringsToInt[extracted.String()]
		if exists {
			return value, exists
		}
	}

	return -1, false
}

func getDigits(s string) []int {
	firstDigit := -1
	lastDigit := -1

	for i, r := range s {
		if !unicode.IsDigit(r) {
			digitFromStr, exists := getDigitsAsString(s[i:])
			if !exists {
				continue
			}

			r = rune('0' + digitFromStr)
		}

		if firstDigit == -1 {
			firstDigit, _ = strconv.Atoi(string(r))
			continue
		}

		if unicode.IsDigit(r) {
			lastDigit, _ = strconv.Atoi(string(r))
		}
	}

	if firstDigit == -1 {
		return []int{}
	} else if lastDigit == -1 {
		return []int{firstDigit, firstDigit}
	} else {
		return []int{firstDigit, lastDigit}
	}
}
