package main

import (
	"gopuzzle/gameMap"
	"gopuzzle/stringPreprocessing"
)


func FindUniqueLocations(directionsStr string) int {

	directionRunes := stringPreprocessing.PreprocessString(&directionsStr)

	inputStringLength := len(directionRunes)

	if inputStringLength < 2 {
		return inputStringLength
	}

	activeMemorySlice := make([]uint64, inputStringLength, inputStringLength)

	currentLocation := gameMap.GenerateOrigin()

	stepNo := 0

	for _, direction := range directionRunes {

		switch direction {
			case 'N':
				gameMap.MoveNorth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
			case 'S':
				gameMap.MoveSouth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
			case 'E':
				gameMap.MoveEast(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
			case 'W':
				gameMap.MoveWest(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
		}
		stepNo++
	}

	return gameMap.CountLocations(&activeMemorySlice, inputStringLength)
}

func main() {
}
