package main

import (
	"fmt"
	"gopuzzle/gameMap"
	"gopuzzle/inputGeneration"
	"time"
)

const MemSize = 66000
type OrdinatesVisitedType [MemSize]uint64
var baseMemory = OrdinatesVisitedType{}

func FindUniqueLocations(directionsStr string) int {

	inputStringLength := len(directionsStr)

	if inputStringLength < 2 {
		return inputStringLength
	}

	activeMemorySlice := (baseMemory)[0:inputStringLength]

	currentLocation := gameMap.GenerateOrigin()

	for stepNo, direction := range directionsStr {

		switch direction {
			case 'N': gameMap.MoveNorth(&currentLocation)
			case 'S': gameMap.MoveSouth(&currentLocation)
			case 'E': gameMap.MoveEast(&currentLocation)
			case 'W': gameMap.MoveWest(&currentLocation)
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
