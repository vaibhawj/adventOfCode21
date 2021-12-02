package main

import (
	"strconv"
	"strings"
)

func Part1(moves []string) int {
	hPos := 0
	depth := 0
	for i := 0; i < len(moves); i++ {
		move := moves[i]
		dir := strings.Split(move, " ")[0]
		x, err := strconv.Atoi(strings.Split(move, " ")[1])
		if err != nil {
			return -1
		}
		switch dir {
		case "forward":
			hPos += x
		case "up":
			depth -= x
		case "down":
			depth += x
		}
	}
	return hPos * depth
}

func Part2(moves []string) int {
	hpos := 0
	depth := 0
	aim := 0
	for i := 0; i < len(moves); i++ {
		move := moves[i]
		dir := strings.Split(move, " ")[0]
		x, err := strconv.Atoi(strings.Split(move, " ")[1])
		if err != nil {
			return -1
		}
		switch dir {
		case "forward":
			hpos += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		}
	}
	return hpos * depth
}
