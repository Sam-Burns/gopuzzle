package inputGeneration

import (
	"strings"
	"testing"
)

func init() {
}

func TestRandomJourney(t *testing.T) {

	journey := RandomJourney(100)

	if len(journey) != 100 {
		t.Fatalf("Random journey was wrong length")
	}

	if !(strings.Contains(journey, "N") && strings.Contains(journey, "S") && strings.Contains(journey, "E") && strings.Contains(journey, "W")) {
		t.Fatalf("Random journeys don't seem to contain everything they should")
	}
}
