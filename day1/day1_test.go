package day1

import "testing"

func TestPart1(t *testing.T) {
	actual := Part1([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	expected := 7
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	expected := 5
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
