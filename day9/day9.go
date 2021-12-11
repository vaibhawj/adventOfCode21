package day9

import (
	"sort"
	"strconv"
	"strings"
)

func RiskLevelsOfLowPoints(input string) int {
	lines := strings.Split(input, "\n")
	heightMap := parseInput(lines)

	lowPoints := getLowPoints(heightMap)
	//fmt.Println(input)
	//fmt.Println(lowPoints)
	sum := 0
	for _, v := range lowPoints {
		sum += v.n + 1
	}

	return sum
}

func ThreeLargestBasins(input string) int {
	lines := strings.Split(input, "\n")
	heightMap := parseInput(lines)
	lowPoints := getLowPoints(heightMap)
	//fmt.Println(input)
	//fmt.Println(lowPoints)
	var basins []basin
	for _, p := range lowPoints {
		var b []point
		scanBasin(p, &b, heightMap)
		basins = append(basins, basin{b})
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].points) > len(basins[j].points)
	})
	result := 1
	for i := 0; i < 3; i++ {
		result *= len(basins[i].points)
	}

	return result
}

func scanBasin(p point, basin *[]point, heightMap [][]int) {
	*basin = append(*basin, p)
	// left
	if p.j > 0 && !contains(point{i: p.i, j: p.j - 1}, basin) && heightMap[p.i][p.j-1] < 9 {
		scanBasin(point{p.i, p.j - 1, heightMap[p.i][p.j-1]}, basin, heightMap)
	}

	// right
	if p.j < len(heightMap[0])-1 && !contains(point{i: p.i, j: p.j + 1}, basin) && heightMap[p.i][p.j+1] < 9 {
		scanBasin(point{p.i, p.j + 1, heightMap[p.i][p.j+1]}, basin, heightMap)
	}

	// up
	if p.i > 0 && !contains(point{i: p.i - 1, j: p.j}, basin) && heightMap[p.i-1][p.j] < 9 {
		scanBasin(point{p.i - 1, p.j, heightMap[p.i-1][p.j]}, basin, heightMap)
	}

	// down
	if p.i < len(heightMap)-1 && !contains(point{i: p.i + 1, j: p.j}, basin) && heightMap[p.i+1][p.j] < 9 {
		scanBasin(point{p.i + 1, p.j, heightMap[p.i+1][p.j]}, basin, heightMap)
	}
}

func contains(p point, scannedPoint *[]point) bool {
	for _, sp := range *scannedPoint {
		if p.i == sp.i && p.j == sp.j {
			return true
		}
	}
	return false
}

func getLowPoints(heightMap [][]int) []point {
	rows := len(heightMap)
	cols := len(heightMap[0])
	var lowPoints []point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			n := heightMap[i][j]
			// left
			if j > 0 && n >= heightMap[i][j-1] {
				continue
			}
			// right
			if j < cols-1 && n >= heightMap[i][j+1] {
				continue
			}
			// up
			if i > 0 && n >= heightMap[i-1][j] {
				continue
			}
			// down
			if i < rows-1 && n >= heightMap[i+1][j] {
				continue
			}
			lowPoints = append(lowPoints, point{i, j, n})
		}
	}
	return lowPoints
}

func parseInput(lines []string) [][]int {
	var heightMap [][]int
	for _, line := range lines {
		var row []int
		for _, s := range strings.Split(line, "") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		heightMap = append(heightMap, row)
	}
	return heightMap
}

type point struct {
	i, j, n int
}

type basin struct {
	points []point
}
