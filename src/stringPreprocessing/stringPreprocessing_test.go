package stringPreprocessing

import (
	"testing"
)

func init() {
}

func TestEmptyString(t *testing.T) {
	testCase := ""
	expectedResult := make([]rune, 0)
	actualResult, _ := PreprocessString(&testCase)
	if !slicesAreEqual(expectedResult, actualResult) {
		t.Fatalf("Failed counting an empty journey: expected %d, got %d", expectedResult, actualResult)
	}
}

func TestLongerString(t *testing.T) {
	testCase := "NWS"
	expectedResult := make([]rune, 0)
	expectedResult = append(expectedResult, 'N', 'W', 'S')
	actualResult, _ := PreprocessString(&testCase)
	if !slicesAreEqual(expectedResult, actualResult) {
		t.Fatalf("Failed counting a non-empty journey: expected %d, got %d", expectedResult, actualResult)
	}
}

func TestStrippingSimpleBackstep(t *testing.T) {
	testCase := "NNSN"
	expectedResult := make([]rune, 0)
	expectedResult = append(expectedResult, 'N', 'N', 'X', 'X')
	actualResult, _ := PreprocessString(&testCase)
	if !slicesAreEqual(expectedResult, actualResult) {
		t.Fatalf("Failed counting a journey with basic backstep stripping: expected %d, got %d", expectedResult, actualResult)
	}
}

func TestMoreStuff(t *testing.T) {
	testCase := "ENSNEWESNSWEW"
	expectedResult := make([]rune, 0)
	expectedResult = append(expectedResult, 'E', 'N', 'X', 'X', 'E', 'X', 'X', 'S', 'X', 'X', 'W', 'X', 'X')
	actualResult, _ := PreprocessString(&testCase)
	if !slicesAreEqual(expectedResult, actualResult) {
		t.Fatalf("Failed counting a journey with multiple backstep strips: expected %d, got %d", expectedResult, actualResult)
	}
}



func slicesAreEqual(a, b []rune) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
