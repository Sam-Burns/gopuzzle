package main

import (
	"fmt"
	"gopuzzle/gameMap"
	"gopuzzle/inputGeneration"
	"time"
)

var locationsVisited = gameMap.OrdinatesVisitedType{}

func FindUniqueLocations(directionsStr string) int {

	currentLocation := gameMap.GenerateOrigin()

	for stepNo, direction := range directionsStr {

		if direction == 'N' {
			currentLocation.MoveNorth()
		} else if direction == 'S' {
			currentLocation.MoveSouth()
		} else if direction == 'E' {
			currentLocation.MoveEast()
		} else if direction == 'W' {
			currentLocation.MoveWest()
		}

		locationsVisited[stepNo] = currentLocation
	}

	return gameMap.CountLocations(&locationsVisited)
}

func main() {

	inputSize := 25000000
	journey := inputGeneration.RandomJourney(inputSize)
	startTime := time.Now()
	uniqueLocations := FindUniqueLocations(journey)
	endTime := time.Now()
	totalTimeNanosecs := endTime.Sub(startTime)

	println("Size of input:    ", inputSize)
	println("Unique locations: ", uniqueLocations)
	print("Total time:        ")
	fmt.Println(totalTimeNanosecs)
}
