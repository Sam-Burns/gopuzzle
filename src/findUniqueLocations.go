package main

import (
	"fmt"
	"gopuzzle/gameMap"
	"gopuzzle/inputGeneration"
	"time"
)

func FindUniqueLocations(directionsStr string) int {

	actualGameMap := gameMap.NewGameMap()

	currentLocation := gameMap.GridLoc{0, 0}

	for _, direction := range directionsStr {

		if direction == 'N' {
			currentLocation.Y += 1
		} else if direction == 'S' {
			currentLocation.Y -= 1
		} else if direction == 'E' {
			currentLocation.X += 1
		} else if direction == 'W' {
			currentLocation.X -= 1
		}

		gameMap.Visit(&currentLocation, &actualGameMap)
	}

	return gameMap.CountPlacesVisited(&actualGameMap)
}

func main() {

	inputSize := 25000000
	journey := inputGeneration.RandomJourney(inputSize)
	startTime := time.Now()
	uniqueLocations := FindUniqueLocations(journey)
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	println("Size of input:    ", inputSize)
	println("Unique locations: ", uniqueLocations)
	print("Total time:        ")
	fmt.Println(totalTime)
}
