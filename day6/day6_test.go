package main

import "testing"

func TestNumberOfLanternFish_P1_WithDemoInput_18Days(t *testing.T) {
	actual := NumberOfLanternFish([]int{3, 4, 3, 1, 2}, 18)
	expected := 26
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumberOfLanternFish_P1_WithDemoInput_80Days(t *testing.T) {
	actual := NumberOfLanternFish([]int{3, 4, 3, 1, 2}, 80)
	expected := 5934
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumberOfLanternFish_P1_WithDemoInput_256Days(t *testing.T) {
	actual := NumberOfLanternFish([]int{3, 4, 3, 1, 2}, 256)
	expected := 26984457539
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumberOfLanternFish_P1_WithRealInput_80Days(t *testing.T) {
	actual := NumberOfLanternFish([]int{4, 3, 3, 5, 4, 1, 2, 1, 3, 1, 1, 1, 1, 1, 2, 4, 1, 3, 3, 1, 1, 1, 1, 2, 3, 1, 1, 1, 4, 1, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 4, 2, 1, 1, 2, 1, 3, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 3, 2, 2, 3, 1, 1, 1, 4, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 1, 3, 1, 1, 2, 4, 1, 1, 3, 2, 4, 1, 1, 1, 1, 1, 5, 5, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 3, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 5, 4, 1, 5, 1, 3, 4, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 2, 1, 2, 3, 5, 1, 1, 1, 1, 3, 5, 1, 1, 1, 2, 1, 1, 4, 1, 1, 5, 1, 4, 1, 2, 1, 3, 1, 5, 1, 4, 3, 1, 3, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1, 4, 5, 1, 1, 1, 1, 1, 3, 1, 3, 4, 1, 1, 4, 1, 1, 3, 1, 3, 1, 1, 4, 5, 4, 3, 2, 5, 1, 1, 1, 1, 1, 1, 2, 1, 5, 2, 5, 3, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 5, 1, 2, 1, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 1, 1, 5, 1, 3, 5, 5, 1, 1, 1, 2, 1, 2, 1, 5, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1}, 80)
	expected := 387413
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestNumberOfLanternFish_P1_WithRealInput_256Days(t *testing.T) {
	actual := NumberOfLanternFish([]int{4, 3, 3, 5, 4, 1, 2, 1, 3, 1, 1, 1, 1, 1, 2, 4, 1, 3, 3, 1, 1, 1, 1, 2, 3, 1, 1, 1, 4, 1, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 4, 2, 1, 1, 2, 1, 3, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 3, 2, 2, 3, 1, 1, 1, 4, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 1, 3, 1, 1, 2, 4, 1, 1, 3, 2, 4, 1, 1, 1, 1, 1, 5, 5, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 3, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 5, 4, 1, 5, 1, 3, 4, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 2, 1, 2, 3, 5, 1, 1, 1, 1, 3, 5, 1, 1, 1, 2, 1, 1, 4, 1, 1, 5, 1, 4, 1, 2, 1, 3, 1, 5, 1, 4, 3, 1, 3, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1, 4, 5, 1, 1, 1, 1, 1, 3, 1, 3, 4, 1, 1, 4, 1, 1, 3, 1, 3, 1, 1, 4, 5, 4, 3, 2, 5, 1, 1, 1, 1, 1, 1, 2, 1, 5, 2, 5, 3, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 5, 1, 2, 1, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 1, 1, 5, 1, 3, 5, 5, 1, 1, 1, 2, 1, 2, 1, 5, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1}, 256)
	expected := 1738377086345
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
