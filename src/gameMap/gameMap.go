package gameMap

var gameMapByteModificationMask = [8]uint8{128, 64, 32, 16, 8, 4, 2, 1}

type GameMapRowType map[uint16]uint8
type GameMapType map[uint16]GameMapRowType

type GridLoc struct {
	X uint16
	Y uint16
}

var xAxisLength int

func NewGameMap(inputString string) GameMapType {

	yAxisLength := len(inputString) / 20
	xAxisLength = yAxisLength / 8

	//return make(map[uint16]GameMapRowType)
	return make(map[uint16]GameMapRowType, yAxisLength)
}

func SetBitNoToTrue(originalByte *uint8, bitNo *int) uint8 {
	return *originalByte | gameMapByteModificationMask[*bitNo]
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

	for _, thisGameRow := range *gameMapToCount {
		for _, thisGameMapByte := range thisGameRow {
			total += CountOnesInByte(&thisGameMapByte)
		}
	}

	return total
}

func Visit (placeVisited *GridLoc, actualGameMap *GameMapType) {

	xByteNo := placeVisited.X / 8

	xBitNo := placeVisited.X % 8

	yRowNo := placeVisited.Y

	_, yValueExists := (*actualGameMap)[yRowNo]

	if !yValueExists {
		(*actualGameMap)[yRowNo] = make(map[uint16]uint8, xAxisLength)
	}

	(*actualGameMap)[yRowNo][xByteNo] = (*actualGameMap)[yRowNo][xByteNo] | gameMapByteModificationMask[xBitNo]
}
