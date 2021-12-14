package day5

import (
	"math"
	"strconv"
	"strings"
)

// part1
func NumberOfPointsWithAtLeastTwoOverlapsP1(input string) int {
	lines, maxPoint := parseInput(input)
	numberOfPoints := 0
	for y := 0; y <= maxPoint; y++ {
		for x := 0; x <= maxPoint; x++ {
			linesCount := 0
			for _, line := range lines {
				if line.x1 == line.x2 && x == line.x1 {
					if y >= int(math.Min(float64(line.y1), float64(line.y2))) && y <= int(math.Max(float64(line.y1), float64(line.y2))) {
						linesCount++
					}
				} else if line.y1 == line.y2 && y == line.y1 {
					if x >= int(math.Min(float64(line.x1), float64(line.x2))) && x <= int(math.Max(float64(line.x1), float64(line.x2))) {
						linesCount++
					}
				}
			}
			if linesCount >= 2 {
				numberOfPoints++
			}
		}
	}

	return numberOfPoints
}

// part2
func NumberOfPointsWithAtLeastTwoOverlapsP2(input string) int {
	lines, maxPoint := parseInput(input)
	numberOfPoints := 0
	for y := 0; y <= maxPoint; y++ {
		for x := 0; x <= maxPoint; x++ {
			linesCount := 0
			for _, line := range lines {
				if line.x1 == line.x2 {
					if x == line.x1 && y >= int(math.Min(float64(line.y1), float64(line.y2))) &&
						y <= int(math.Max(float64(line.y1), float64(line.y2))) {
						linesCount++
					}
				} else if line.y1 == line.y2 {
					if y == line.y1 && x >= int(math.Min(float64(line.x1), float64(line.x2))) &&
						x <= int(math.Max(float64(line.x1), float64(line.x2))) {
						linesCount++
					}
				} else {
					if x >= int(math.Min(float64(line.x1), float64(line.x2))) &&
						x <= int(math.Max(float64(line.x1), float64(line.x2))) &&
						y >= int(math.Min(float64(line.y1), float64(line.y2))) &&
						y <= int(math.Max(float64(line.y1), float64(line.y2))) &&
						math.Abs(float64(y-line.y1)) == math.Abs(float64(x-line.x1)) {
						linesCount++
					}
				}
			}
			if linesCount >= 2 {
				numberOfPoints++
			}
			//fmt.Printf("%d ", linesCount)
		}
		//fmt.Println()
	}

	return numberOfPoints
}

func parseInput(input string) ([]Line, int) {
	inputLines := strings.Split(input, "\n")
	maxX := float64(0)
	maxY := float64(0)
	var lines []Line
	for _, line := range inputLines {
		endPoints := strings.Split(line, " -> ")
		x1, _ := strconv.ParseInt(strings.Split(endPoints[0], ",")[0], 10, 64)
		y1, _ := strconv.ParseInt(strings.Split(endPoints[0], ",")[1], 10, 64)

		x2, _ := strconv.ParseInt(strings.Split(endPoints[1], ",")[0], 10, 64)
		y2, _ := strconv.ParseInt(strings.Split(endPoints[1], ",")[1], 10, 64)

		maxOfX1AndX2 := math.Max(float64(x1), float64(x2))
		if maxOfX1AndX2 > maxX {
			maxX = maxOfX1AndX2
		}
		maxOfY1AndY2 := math.Max(float64(y1), float64(y2))
		if maxOfY1AndY2 > maxY {
			maxY = maxOfX1AndX2
		}
		lines = append(lines, Line{x1: int(x1), y1: int(y1), x2: int(x2), y2: int(y2)})
	}
	maxPoint := int(math.Max(maxX, maxY))
	return lines, maxPoint
}

type Line struct {
	x1, y1, x2, y2 int
}
