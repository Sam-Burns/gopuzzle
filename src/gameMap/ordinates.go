package gameMap

type OrdinatesVisitedType [50000]uint64

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

func (ordinatesVisited *OrdinatesVisitedType) Clobber() {
	for index, _ := range *ordinatesVisited {
		(*ordinatesVisited)[index] = 0
	}
}
