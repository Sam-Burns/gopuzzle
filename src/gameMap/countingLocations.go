package gameMap

import (
	"sort"
)

func CountLocations(ordinatesVisited *[]uint64, inputStringLength int) int {

	SortArray(ordinatesVisited)

	return countUniqueOrdinates(ordinatesVisited, &inputStringLength)
}

func SortArray(ordinatesVisited *[]uint64) {
	sort.Slice(
		*ordinatesVisited,
		func(i, j int) bool {
			return (*ordinatesVisited)[i] < (*ordinatesVisited)[j]
		})
}

func countUniqueOrdinates(ordinatesVisited *[]uint64, inputStringLength *int) int {
	count := 1
	index := 1

	for index <= *inputStringLength - 1 {
		if (*ordinatesVisited)[index] != 0 && (*ordinatesVisited)[index] != (*ordinatesVisited)[index-1] {
			count++
		}
		index++
	}

	return count
}
