package gameMap

import (
	"sort"
)

func CountLocations(ordinatesVisited *OrdinatesVisitedType, inputStringLength int) int {

	sortArray(ordinatesVisited, &inputStringLength)

	return countUniqueOrdinates(ordinatesVisited, &inputStringLength)
}

func sortArray(ordinatesVisited *OrdinatesVisitedType, inputStringLength *int) {
	mySlice := (*ordinatesVisited)[0:*inputStringLength]

	sort.Slice(
		mySlice,
		func(i, j int) bool {
			return (*ordinatesVisited)[i] < (*ordinatesVisited)[j]
		})
}

func countUniqueOrdinates(ordinatesVisited *OrdinatesVisitedType, inputStringLength *int) int {
	count := 1
	index := 1

	for index <= *inputStringLength - 1 {
		if (*ordinatesVisited)[index] != (*ordinatesVisited)[index-1] {
			count++
		}
		index++
	}

	return count
}
