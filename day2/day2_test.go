package day2

import "testing"

func TestPart1(t *testing.T) {
	actual := Part1([]string{"forward 5","down 5","forward 8","up 3","down 8","forward 2"})
	expected := 150
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2([]string{"forward 5","down 5","forward 8","up 3","down 8","forward 2"})
	expected := 900
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
