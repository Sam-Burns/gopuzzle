package gameMap

const binary1OnXAxis uint64 = 4294967296 // 1 << 32
const binaryOrigin uint64 = 9223372039002259456 // (1 << 63) + (1 << 31)

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

func MoveSoutheast(location *uint64) {
	MoveEast(location)
	MoveSouth(location)
}

func MoveNortheast(location *uint64) {
	MoveEast(location)
	MoveNorth(location)
}

func MoveSouthwest(location *uint64) {
	MoveWest(location)
	MoveSouth(location)
}

func MoveNorthwest(location *uint64) {
	MoveWest(location)
	MoveNorth(location)
}
