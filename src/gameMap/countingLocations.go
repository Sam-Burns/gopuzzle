package gameMap

import (
	"bytes"
	"encoding/binary"
)

func CountLocations(ordinatesVisited *OrdinatesVisitedType) int {

	count := 0

	SortOrdinates(ordinatesVisited)

	for index, value := range ordinatesVisited {
		if index == 0 {
			count++
		} else if value == 0 {
		} else if value != ordinatesVisited[index-1] {
			count++
		}
	}

	return count
}

func SortOrdinates(ordinatesToSort *OrdinatesVisitedType) {
	radixsort(ordinatesToSort)
}

const digit = 4
const maxbit uint64 = 1 << 63

func radixsort(data *OrdinatesVisitedType) {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(*data))
	for i, e := range *data {
		_ = binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		_, _ = buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w uint64
	for i, b := range ds {
		buf.Write(b)
		_ = binary.Read(buf, binary.LittleEndian, &w)
		(*data)[i] = w^maxbit
	}
}
