package main

import (
	"testing"
)

func init() {
}

func TestASingleMovementCanBeCounted(t *testing.T) {
	testCase := "N"
	expectedResult := 1
	result := FindUniqueLocations(testCase)
	if result != expectedResult {
		t.Fatalf("Failed counting a single move: expected %d, got %d", expectedResult, result)
	}
}

func TestMultipleMovementsCanBeCounted(t *testing.T) {
	testCase := "NE"
	expectedResult := 2
	result := FindUniqueLocations(testCase)
	if result != expectedResult {
		t.Fatalf("Failed counting a single move: expected %d, got %d", expectedResult, result)
	}
}
