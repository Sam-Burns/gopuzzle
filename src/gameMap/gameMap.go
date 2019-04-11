package gameMap

var gameMapByteModificationMask = [8]uint8{128, 64, 32, 16, 8, 4, 2, 1}

func SetBitNoToTrue(originalByte uint8, bitNo int) uint8 {
	return originalByte | gameMapByteModificationMask[bitNo]
}

func CountOnesInByte(byteToCount uint8) int {

	total := 0

	for bitNo := uint8(0); bitNo <= 7; bitNo++ {

		shiftedThing := byteToCount >> bitNo

		if ((shiftedThing) % 2) == 1 {
			total += 1
		}
	}

	return total
}
