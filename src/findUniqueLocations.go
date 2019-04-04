package main

import (
	"fmt"
	"time"
)

type gridLoc struct {
	x int
	y int
}

var placesVisited = make(map[int]map[int]bool)

func FindUniqueLocations(directionsStr string) int {

	currentLocation := gridLoc{0, 0}

	for _, direction := range directionsStr {

		if direction == 'N' {
			currentLocation = gridLoc{currentLocation.x, currentLocation.y + 1}
		} else if direction == 'S' {
			currentLocation = gridLoc{currentLocation.x, currentLocation.y - 1}
		} else if direction == 'E' {
			currentLocation = gridLoc{currentLocation.x + 1, currentLocation.y}
		} else if direction == 'W' {
			currentLocation = gridLoc{currentLocation.x - 1, currentLocation.y}
		}

		visit(currentLocation)
	}

	return len(placesVisited)
}

func visit (placeVisited gridLoc) {
	_, alreadyVisitedXOrdinate := placesVisited[placeVisited.x]
	if ! alreadyVisitedXOrdinate {
		placesVisited[placeVisited.x] = make(map[int]bool)
	}
	placesVisited[placeVisited.x][placeVisited.y] = true
}

func main() {

	inputSize := 100
	journey := RandomJourney(inputSize)
	startTime := time.Now()
	uniqueLocations := FindUniqueLocations(journey)
	endTime := time.Now()
	totalTimeNanosecs := endTime.Sub(startTime)

	println("Size of input:    ", inputSize)
	println("Unique locations: ", uniqueLocations)
	print("Total time:        ")
	fmt.Println(totalTimeNanosecs)
}
