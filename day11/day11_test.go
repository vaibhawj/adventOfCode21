package day11

import (
	"testing"
)

func TestTotalFlashes(t *testing.T) {
	input := "5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"
	actual := TotalFlashes(input, 100)
	expected := 1656
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestTotalFlashes_Real(t *testing.T) {
	input := "8271653836\n7567626775\n2315713316\n6542655315\n2453637333\n1247264328\n2325146614\n2115843171\n6182376282\n2384738675"
	actual := TotalFlashes(input, 100)
	expected := 1562
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestSyncStep(t *testing.T) {
	input := "5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"
	actual := SyncStep(input)
	expected := 195
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestSyncStep_Real(t *testing.T) {
	input := "8271653836\n7567626775\n2315713316\n6542655315\n2453637333\n1247264328\n2325146614\n2115843171\n6182376282\n2384738675"
	actual := SyncStep(input)
	expected := 268
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}