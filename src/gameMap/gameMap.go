package gameMap

var gameMapByteModificationMask = [8]uint8{128, 64, 32, 16, 8, 4, 2, 1}

type GameMapRowType [2048]uint8
type GameMapType [16384]GameMapRowType

type GridLoc struct {
	X int
	Y int
}

func NewGameMap() GameMapType {
	var actualGameMap GameMapType
	return actualGameMap
}

func SetBitNoToTrue(originalByte uint8, bitNo int) uint8 {
	return originalByte | gameMapByteModificationMask[bitNo]
}

func CountOnesInByte(byteToCount *uint8) int {

	total := 0

	for bitNo := uint8(0); bitNo <= 7; bitNo++ {

		shiftedThing := *byteToCount >> bitNo

		if ((shiftedThing) % 2) == 1 {
			total += 1
		}
	}

	return total
}

func CountPlacesVisited(gameMapToCount *GameMapType) int {

	total := 0

	for _, thisGameRow := range gameMapToCount {
		for _, thisGameMapByte := range thisGameRow {
			total += CountOnesInByte(&thisGameMapByte)
		}
	}

	return total
}

func Visit (placeVisited *GridLoc, actualGameMap *GameMapType) {

	xByteNo := getByteNoOnXAxis(placeVisited.X)

	xBitNo := getBitNoOnXAxis(placeVisited.X)

	yRowNo := getRowNoOnYAxis(placeVisited.Y)

	actualGameMap[yRowNo][xByteNo] = SetBitNoToTrue(actualGameMap[yRowNo][xByteNo], xBitNo)
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
