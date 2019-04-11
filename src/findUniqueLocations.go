package main

import (
	"fmt"
	"gopuzzle/gameMap"
	"gopuzzle/inputGeneration"
	//"gopuzzle/reliableResult"
	"time"
)

type gridLoc struct {
	x int
	y int
}

var placesVisited = make(map[int]map[int]bool)

type gameMapRowType [2048]uint8
type gameMapType [16384]gameMapRowType

var actualGameMap gameMapType

func FindUniqueLocations(directionsStr string) int {

	currentLocation := gridLoc{8192, 8192}

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

	return countPlacesVisited(actualGameMap)
}

func visit (placeVisited gridLoc) {

	xByteNo := getByteNoOnXAxis(placeVisited.x)

	xBitNo := getBitNoOnXAxis(placeVisited.x)

	yRowNo := getRowNoOnYAxis(placeVisited.y)


	actualGameMap[yRowNo][xByteNo] = gameMap.SetBitNoToTrue(actualGameMap[yRowNo][xByteNo], xBitNo)

	_, alreadyVisitedXOrdinate := placesVisited[placeVisited.x]
	if ! alreadyVisitedXOrdinate {
		placesVisited[placeVisited.x] = make(map[int]bool)
	}
	placesVisited[placeVisited.x][placeVisited.y] = true
}

func getByteNoOnXAxis(ordinate int) int {
	return ordinate / 8
}

func getBitNoOnXAxis(ordinate int) int {
	return ordinate % 8
}

func getRowNoOnYAxis(ordinate int) int {
	return ordinate
}

func countPlacesVisited(gameMapToCount gameMapType) int {

	total := 0

	for _, thisGameRow := range gameMapToCount {
		for _, thisGameMapByte := range thisGameRow {
			total += gameMap.CountOnesInByte(thisGameMapByte)
		}
	}

	return total
}

func main() {

	inputSize := 25000000
	//inputSize := 5000000
	journey := inputGeneration.RandomJourney(inputSize)
	startTime := time.Now()
	uniqueLocations := FindUniqueLocations(journey)
	endTime := time.Now()
	totalTimeNanosecs := endTime.Sub(startTime)
	//reliableUniqueLocations := reliableResult.FindUniqueLocations(journey)

	println("Size of input:    ", inputSize)
	//println("Input string:     ", journey)
	println("Unique locations: ", uniqueLocations)
	//println("Reliable result:  ", reliableUniqueLocations)
	print("Total time:        ")
	fmt.Println(totalTimeNanosecs)
}
