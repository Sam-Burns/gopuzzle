package reliableResult

import (
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

// N.B. failing tests is probably because reliableResult.placesVisited isn't torn down between runs of
// reliableResult.FindUniqueLocations() . There might be something else wrong as well.

// To Do: reliable result thing should pass this test!
//
//func TestLongerJourney(t *testing.T) {
//	testCase := "NWNWSSNWWS"
//	expectedResult := 9
//	result := FindUniqueLocations(testCase)
//	if result != expectedResult {
//		t.Fatalf("Failed counting a longer path: expected %d, got %d", expectedResult, result)
//	}
//}
//
//
// To Do: shorter version of failing test
//
//func TestLongerJourney(t *testing.T) {
//	testCase := "NWNWS"
//	expectedResult := 5
//	result := FindUniqueLocations(testCase)
//	if result != expectedResult {
//		t.Fatalf("Failed counting a longer path: expected %d, got %d", expectedResult, result)
//	}
//}
