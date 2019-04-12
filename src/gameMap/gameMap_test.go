package gameMap

import "testing"

func init() {
}

func TestItCanDoSomething(t *testing.T) {

	result := SetBitNoToTrue(0, 0)
	if result != 128 {
		t.Fatalf("Failed setting left-hand bit, got %d", result)
	}

	otherResult := SetBitNoToTrue(0, 4)
	if otherResult != 8 {
		t.Fatalf("Failed setting a bit, got %d", otherResult)
	}
}

func TestItCanCountOnesInByte(t *testing.T) {

	var testCase1 uint8
	testCase1 = 16
	result1 := CountOnesInByte(&testCase1)

	if result1 != 1 {
		t.Fatalf("Failed counting bits, got %d", result1)
	}

	var testCase2 uint8
	testCase2 = 33
	result2 := CountOnesInByte(&testCase2)

	if result2 != 2 {
		t.Fatalf("Failed counting bits, got %d", result2)
	}
}

func TestItCanCountPlacesVisited(t *testing.T) {

	var gameMap GameMapType

	result := CountPlacesVisited(&gameMap)

	if result != 0 {
		t.Fatalf("Failed counting places visted, got %d", result)
	}

	Visit(
		&GridLoc{
			X: 1,
			Y: 1,
		},
		&gameMap,
	)

	result2 := CountPlacesVisited(&gameMap)

	if result2 != 1 {
		t.Fatalf("Failed counting places visted, got %d", result2)
	}
}
