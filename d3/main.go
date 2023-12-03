package main

import (
	"github.com/alc6/aoc2023/util"
	"log"
	"maps"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	lines, err := util.ReadFileLines("d3/input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	schematic2D := make([][]string, len(lines))
	for i, y := range lines {
		for _, x := range y {
			schematic2D[i] = append(schematic2D[i], string(x))
		}
	}

	var (
		gatheredPoints          = findDigitPoints(schematic2D)
		pointsWithNeighborhoods = make(PointsWithNeighborhoods, 0, len(gatheredPoints))
	)
	for _, gp := range gatheredPoints {
		pwn := PointsWithNeighborhood{Points: gp}
		pwn.FindNeighbors(schematic2D)
		pointsWithNeighborhoods = append(pointsWithNeighborhoods, pwn)
	}

	pt1(pointsWithNeighborhoods)
	pt2(pointsWithNeighborhoods)
}

func pt1(pointsWithNeighborhoods PointsWithNeighborhoods) {
	var sum int
	for _, pwn := range pointsWithNeighborhoods {
		if pwn.HasSymbolNeighbor() {
			sum += pwn.Points.GetValue()
		}
	}

	log.Println("pt1 sum: ", sum)
}

func pt2(pointsWithNeighborhoods PointsWithNeighborhoods) {
	possibleGears := pointsWithNeighborhoods.FindCommonNeighborsWithSymbol("*")
	maps.DeleteFunc(possibleGears, func(key string, value []PointsWithNeighborhood) bool {
		return len(value) != 2
	})

	var sumGearRatios int
	for _, values := range possibleGears {
		gearRatio := 1
		for numbers := range values {
			gearRatio *= values[numbers].Points.GetValue()
		}
		sumGearRatios += gearRatio
	}

	log.Println("pt2 sum: ", sumGearRatios)
}

type PointsWithNeighborhood struct {
	Points       Points
	Neighborhood Points
}

func (pwn *PointsWithNeighborhood) FindNeighbors(schematic2D [][]string) {
	for _, point := range pwn.Points {
		neighbors := []Point{
			{X: point.X - 1, Y: point.Y - 1},
			{X: point.X - 1, Y: point.Y},
			{X: point.X - 1, Y: point.Y + 1},

			{X: point.X + 1, Y: point.Y - 1},
			{X: point.X + 1, Y: point.Y},
			{X: point.X + 1, Y: point.Y + 1},

			{X: point.X, Y: point.Y - 1},
			{X: point.X, Y: point.Y + 1},
		}
		// Delete points out of layout from those we previously ventilated
		neighbors = slices.DeleteFunc(neighbors, func(p Point) bool {
			return p.X < 0 || p.Y < 0 || p.X >= len(schematic2D[0]) || p.Y >= len(schematic2D)
		})

		for j, neighbor := range neighbors {
			neighbors[j].Value = schematic2D[neighbor.Y][neighbor.X]
		}

		pwn.Neighborhood = append(pwn.Neighborhood, neighbors...)
	}
}

func (pwn *PointsWithNeighborhood) HasSymbolNeighbor() bool {
	for _, neighbor := range pwn.Neighborhood {
		if r, _ := utf8.DecodeRuneInString(neighbor.Value); !unicode.IsDigit(r) && neighbor.Value != "." {
			return true
		}
	}

	return false
}

type PointsWithNeighborhoods []PointsWithNeighborhood

func (pwns PointsWithNeighborhoods) FindCommonNeighborsWithSymbol(symbol string) map[string][]PointsWithNeighborhood {
	// Isolate points with their neighborhoods that contains a neighbor point with the given symbol
	var symbolsNextTo = make(map[string][]PointsWithNeighborhood)
	for _, pwn := range pwns {
		for _, neighbor := range pwn.Neighborhood {
			if neighbor.Value == symbol {
				symbolsNextTo[neighbor.StringCoords()] = append(symbolsNextTo[neighbor.StringCoords()], pwn)
				break
			}
		}
	}

	return symbolsNextTo
}

type Point struct {
	Value string
	X     int
	Y     int
}

func (p Point) StringCoords() string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(p.X))
	sb.WriteString(",")
	sb.WriteString(strconv.Itoa(p.Y))
	return sb.String()
}

type Points []Point

func (p Points) GetValue() int {
	var sb strings.Builder
	for _, point := range p {
		sb.WriteString(point.Value)
	}

	value, _ := strconv.Atoi(sb.String())
	return value
}

func findDigitPoints(schematic [][]string) [][]Point {
	points := make([]Point, 0)

	for y, line := range schematic {
		for x, char := range line {
			if r, _ := utf8.DecodeRuneInString(char); !unicode.IsDigit(r) {
				continue
			}

			points = append(points, Point{Value: char, X: x, Y: y})
		}
	}

	gatheredPoints := make([][]Point, 0)
	for i, p := range points {
		if i == 0 {
			gatheredPoints = append(gatheredPoints, []Point{p})
			continue
		}

		if p.X == points[i-1].X+1 && p.Y == points[i-1].Y {
			gatheredPoints[len(gatheredPoints)-1] = append(gatheredPoints[len(gatheredPoints)-1], p)
		} else {
			gatheredPoints = append(gatheredPoints, []Point{p})
		}
	}

	return gatheredPoints
}
