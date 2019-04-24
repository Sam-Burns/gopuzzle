package gameMap

//const MemSize = 50000
const MemSize = 66000

type OrdinatesVisitedType [MemSize]uint64

func GenerateOrigin() uint64 {
	return (1 << 63) + (1 << 31)
}

func MoveNorth(location *uint64) {
	*location += 1
}

func MoveSouth(location *uint64) {
	*location -= 1
}

func MoveEast(location *uint64) {
	*location += 1 << 32
}

func MoveWest(location *uint64) {
	*location -= 1 << 32
}

func (ordinatesVisited *OrdinatesVisitedType) Clobber(inputStringLength int) {

	index := 0

	for index < inputStringLength {
		(*ordinatesVisited)[index] = 0
		index++
	}
}

func(ordinatesVisited *OrdinatesVisitedType) getLength() int {
	index := 0
	for index < MemSize {

		if (*ordinatesVisited)[index] == 0 {
			return index + 1
		}

		index++
	}
	return -1
}
