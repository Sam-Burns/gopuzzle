package main

import (
	"gopuzzle/gameMap"
	"gopuzzle/stringPreprocessing"
)


func FindUniqueLocations(directionsStr string) int {

	directionRunes, directionsNullified := stringPreprocessing.PreprocessString(&directionsStr)

	inputStringLength := len(directionRunes) - directionsNullified

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
			case '3':
				gameMap.MoveSoutheast(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '9':
				gameMap.MoveNortheast(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '7':
				gameMap.MoveNorthwest(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '1':
				gameMap.MoveSouthwest(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '8':
				gameMap.MoveDoubleNorth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '2':
				gameMap.MoveDoubleSouth(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '6':
				gameMap.MoveDoubleEast(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
			case '4':
				gameMap.MoveDoubleWest(&currentLocation)
				activeMemorySlice[stepNo] = currentLocation
				stepNo++
		}
	}

	return gameMap.CountLocations(&activeMemorySlice, inputStringLength)
}

func main() {
}
