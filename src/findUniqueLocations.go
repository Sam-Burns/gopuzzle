package main

import (
	"fmt"
	"gopuzzle/gameMap"
	"gopuzzle/inputGeneration"
	"time"
)

var locationsVisited = gameMap.OrdinatesVisitedType{}

func FindUniqueLocations(directionsStr string) int {

	inputStringLength := len(directionsStr)

	if inputStringLength < 2 {
		return inputStringLength
	}

	activeMemorySlice := (locationsVisited)[0:inputStringLength]

	currentLocation := gameMap.GenerateOrigin()

	for stepNo, direction := range directionsStr {

		if direction == 'N' {
			gameMap.MoveNorth(&currentLocation)
		} else if direction == 'S' {
			gameMap.MoveSouth(&currentLocation)
		} else if direction == 'E' {
			gameMap.MoveEast(&currentLocation)
		} else if direction == 'W' {
			gameMap.MoveWest(&currentLocation)
		}

		activeMemorySlice[stepNo] = currentLocation
	}

	return gameMap.CountLocations(&activeMemorySlice, inputStringLength)
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
