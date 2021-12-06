package main

import "fmt"

func WinnerScore(numbers []int, grids [][][]int) int {
	for _, n := range numbers {
		mark(grids, n)
		winnerIndex := checkWinner(grids)
		if winnerIndex > -1 {
			return n * sumOfUnmarked(grids[winnerIndex])
		}
	}

	return -1
}

func LoserScore(numbers []int, grids [][][]int) int {

	loserIndex := -1
	for _, n := range numbers {
		mark(grids, n)
		winnerCounts := 0
		for i, grid := range grids {
			if isWinner(grid) {
				winnerCounts++
			} else {
				loserIndex = i
			}
		}
		if winnerCounts == len(grids) {
			fmt.Println(n)
			loserGrid := grids[loserIndex]
			return sumOfUnmarked(loserGrid) * n
		}
	}

	return -1
}

func isWinner(grid [][]int) bool {
	for i := 0; i < 5; i++ {
		if sumOfRow(grid, i) == -5 || sumOfColumn(grid, i) == -5 {
			return true
		}
	}
	return false
}

func sumOfUnmarked(grid [][]int) int {
	result := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] > -1 {
				result += grid[i][j]
			}
		}
	}
	return result
}

func checkWinner(grids [][][]int) int {
	for gi, grid := range grids {
		for i := 0; i < 5; i++ {
			if sumOfRow(grid, i) == -5 {
				return gi
			}
			if sumOfColumn(grid, i) == -5 {
				return gi
			}

		}
	}
	return -1
}

func sumOfColumn(grid [][]int, index int) int {
	result := 0
	for i := 0; i < 5; i++ {
		result += grid[i][index]
	}
	return result
}

func sumOfRow(grid [][]int, index int) int {
	result := 0
	for _, v := range grid[index] {
		result += v
	}
	return result
}

func mark(grids [][][]int, n int) {
	for _, grid := range grids {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if grid[i][j] == n {
					grid[i][j] = -1
				}
			}
		}
	}
}

func print(grids [][][]int) {
	for _, grid := range grids {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				fmt.Printf("%d ", grid[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}
