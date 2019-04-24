package gameMap

type OrdinateType uint64

func GenerateOrigin() OrdinateType {
	return (1 << 63) + (1 << 31)
}

func (location *OrdinateType) MoveNorth() {
	*location += 1
}

func (location *OrdinateType) MoveSouth() {
	*location -= 1
}

func (location *OrdinateType) MoveEast() {
	*location += 1 << 32
}

func (location *OrdinateType) MoveWest() {
	*location -= 1 << 32
}

type OrdinatesVisitedType [50000]OrdinateType

func (ordinatesVisited *OrdinatesVisitedType) Clobber() {
	for index, _ := range *ordinatesVisited {
		(*ordinatesVisited)[index] = 0
	}
}
