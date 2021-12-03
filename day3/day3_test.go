package main

import "testing"

func TestPowerConsumption(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	actual := PowerConsumption(input)
	expected := int64(198)
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestLifeSupportRating(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	actual := LifeSupportRating(input)
	expected := int64(230)
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
