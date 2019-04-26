package main

import (
	"gopuzzle/gameMap"
	"gopuzzle/stringPreprocessing"
)


func FindUniqueLocations(directionsStr string) int {

	directionRunes, directionsNullified := stringPreprocessing.PreprocessString(&directionsStr)

	inputStringLength := len(directionRunes)
	postProcessingInputStringLength := inputStringLength - directionsNullified

	if postProcessingInputStringLength < 2 {
		return postProcessingInputStringLength
	}

	activeMemorySlice := make([]uint64, postProcessingInputStringLength, postProcessingInputStringLength)

	currentLocation := gameMap.GenerateOrigin()

	stepNo := 0

	for _, direction := range directionRunes {

		switch direction {
			case 'N':
				gameMap.MoveNorth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case 'S':
				gameMap.MoveSouth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case 'E':
				gameMap.MoveEast(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case 'W':
				gameMap.MoveWest(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
		}
	}

	return gameMap.CountLocations(&activeMemorySlice, postProcessingInputStringLength)
}

func main() {
}
