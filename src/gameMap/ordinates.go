package gameMap

//const MemSize = 50000
const MemSize = 66000

const binary1OnXAxis uint64 = 4294967296 // 1 << 32
const binaryOrigin uint64 = 9223372039002259456 // (1 << 63) + (1 << 31)

type OrdinatesVisitedType [MemSize]uint64

func GenerateOrigin() uint64 {
	return binaryOrigin
}

func MoveNorth(location *uint64) {
	*location += 1
}

func MoveSouth(location *uint64) {
	*location -= 1
}

func MoveEast(location *uint64) {
	*location += binary1OnXAxis
}

func MoveWest(location *uint64) {
	*location -= binary1OnXAxis
}

func (ordinatesVisited *OrdinatesVisitedType) Clobber(inputStringLength int) {

	index := 0

	for index < inputStringLength {
		(*ordinatesVisited)[index] = 0
		index++
	}
}
