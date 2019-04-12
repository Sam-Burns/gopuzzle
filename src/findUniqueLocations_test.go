package main

import (
	"fmt"
	"testing"
)

func init() {
}

func TestEmptyString(t *testing.T) {
	testCase := ""
	expectedResult := 0
	result := FindUniqueLocations(testCase)
	if result != expectedResult {
		t.Fatalf("Failed counting an empty journey: expected %d, got %d", expectedResult, result)
	}
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
		t.Fatalf("Failed counting a two moves: expected %d, got %d", expectedResult, result)
	}
}

func TestDuplicateLocationsAreNotCounted(t *testing.T) {
	testCase := "NEW"
	expectedResult := 2
	result := FindUniqueLocations(testCase)
	if result != expectedResult {
		t.Fatalf("Failed counting a path where a location is visited twice: expected %d, got %d", expectedResult, result)
	}
}

func TestDebug(t *testing.T) {
	testCase := "NWNW"
	expectedResult := 4
	result := FindUniqueLocations(testCase)
	if result != expectedResult {

		fmt.Println()

		t.Fatalf("Failed counting a path where a location is visited twice: expected %d, got %d", expectedResult, result)
	}
}

func TestLongerJourney(t *testing.T) {
	testCase := "NWNWSSNWWS"
	expectedResult := 9
	result := FindUniqueLocations(testCase)
	if result != expectedResult {
		t.Fatalf("Failed counting a longer path: expected %d, got %d", expectedResult, result)
	}
}
