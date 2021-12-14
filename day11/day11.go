package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var flashCount = 0

func TotalFlashes(input string, steps int) int {
	grid := parse(input)
	//print(grid)
	for s := 0; s < steps; s++ {

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j].increment()
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j].flashed && !grid[i][j].incrementedOthers {
					incrementAdjacency(i, j, grid)
					grid[i][j].incrementedOthers = true
				}
			}
		}

		resetFlashStatus(grid)

		//fmt.Printf("Step %d\n", s+1)
		//print(grid)
	}

	return flashCount
}

func SyncStep(input string) int {
	grid := parse(input)
	//print(grid)
	for s := 0; s < math.MaxInt64; s++ {

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j].increment()
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j].flashed && !grid[i][j].incrementedOthers {
					incrementAdjacency(i, j, grid)
					grid[i][j].incrementedOthers = true
				}
			}
		}

		resetFlashStatus(grid)

		if isSync(grid) {
			return s + 1
		}

		//fmt.Printf("Step %d\n", s+1)
		//print(grid)
	}
	return -1
}

func isSync(grid [][]cell) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j].v != 0 {
				return false
			}
		}
	}
	return true
}

func resetFlashStatus(grid [][]cell) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j].flashed = false
			grid[i][j].incrementedOthers = false
		}
	}
}

func incrementAdjacency(i int, j int, grid [][]cell) {
	// left
	if j > 0 {
		if !grid[i][j-1].flashed {
			grid[i][j-1].increment()
			if grid[i][j-1].flashed {
				incrementAdjacency(i, j-1, grid)
				grid[i][j-1].incrementedOthers = true
			}
		}
	}
	// top left
	if i > 0 && j > 0 {
		if !grid[i-1][j-1].flashed {
			grid[i-1][j-1].increment()
			if grid[i-1][j-1].flashed {
				incrementAdjacency(i-1, j-1, grid)
				grid[i-1][j-1].incrementedOthers = true
			}
		}
	}
	// up
	if i > 0 {
		if !grid[i-1][j].flashed {
			grid[i-1][j].increment()
			if grid[i-1][j].flashed {
				incrementAdjacency(i-1, j, grid)
				grid[i-1][j].incrementedOthers = true
			}
		}
	}
	// top right
	if i > 0 && j < len(grid[i])-1 {
		if !grid[i-1][j+1].flashed {
			grid[i-1][j+1].increment()
			if grid[i-1][j+1].flashed {
				incrementAdjacency(i-1, j+1, grid)
				grid[i-1][j+1].incrementedOthers = true
			}
		}
	}
	// right
	if j < len(grid[i])-1 {
		if !grid[i][j+1].flashed {
			grid[i][j+1].increment()
			if grid[i][j+1].flashed {
				incrementAdjacency(i, j+1, grid)
				grid[i][j+1].incrementedOthers = true
			}
		}
	}
	// bottom right
	if i < len(grid)-1 && j < len(grid[i])-1 {
		if !grid[i+1][j+1].flashed {
			grid[i+1][j+1].increment()
			if grid[i+1][j+1].flashed {
				incrementAdjacency(i+1, j+1, grid)
				grid[i+1][j+1].incrementedOthers = true
			}
		}
	}
	//down
	if i < len(grid)-1 {
		if !grid[i+1][j].flashed {
			grid[i+1][j].increment()
			if grid[i+1][j].flashed {
				incrementAdjacency(i+1, j, grid)
				grid[i+1][j].incrementedOthers = true
			}
		}
	}
	//bottom left
	if i < len(grid)-1 && j > 0 {
		if !grid[i+1][j-1].flashed {
			grid[i+1][j-1].increment()
			if grid[i+1][j-1].flashed {
				incrementAdjacency(i+1, j-1, grid)
				grid[i+1][j-1].incrementedOthers = true
			}
		}
	}
}

func parse(input string) [][]cell {
	var grid [][]cell
	for i, line := range strings.Split(input, "\n") {
		var row []cell
		for j, s := range strings.Split(line, "") {
			n, _ := strconv.Atoi(s)
			row = append(row, cell{i, j, n, false, false})
		}
		grid = append(grid, row)
	}
	return grid
}

func print(grid [][]cell) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%d ", grid[i][j].v)
		}
		fmt.Println()
	}
	fmt.Println()
}

type cell struct {
	i, j, v                    int
	flashed, incrementedOthers bool
}

func (c *cell) increment() {
	c.v += 1
	if c.v > 9 {
		c.v = 0
		c.flashed = true
		flashCount += 1
	}
}
