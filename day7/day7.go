package day7

import (
	"math"
)

func AlignEfficientlyPart1(positions []int) int {
	minCost := math.MaxInt64
	for _, p := range positions {
		cost := 0
		for _, v := range positions {
			cost += int(math.Abs(float64(p - v)))
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func AlignEfficientlyPart2(positions []int) int {
	minCost := math.MaxInt64
	for i := min(positions); i <= max(positions); i++ {
		cost := 0
		for _, v := range positions {
			cost += costOfMove(int(math.Abs(float64(i - v))))
		}
		//fmt.Printf("cost of %d = %d\n", i, cost)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func min(positions []int) int {
	min := math.MaxInt64
	for _, v := range positions {
		if v < min {
			min = v
		}
	}
	return min
}

func max(positions []int) int {
	max := math.MinInt64
	for _, v := range positions {
		if v > max {
			max = v
		}
	}
	return max
}

func costOfMove(steps int) int {
	sum := 0
	for i := 1; i <= steps; i++ {
		sum += i
	}
	return sum
}
