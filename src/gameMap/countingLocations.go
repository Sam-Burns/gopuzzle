package gameMap

import (
	"math/rand"
)

func CountLocations(ordinatesVisited *[]uint64, inputStringLength int) int {

	quicksort(ordinatesVisited)

	return countUniqueOrdinates(ordinatesVisited, &inputStringLength)
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

func quicksort(sliceToSort *[]uint64) {
	if len(*sliceToSort) < 2 {
		return
	}

	left, right := 0, len(*sliceToSort)-1

	pivot := rand.Int() % len(*sliceToSort)

	(*sliceToSort)[pivot], (*sliceToSort)[right] = (*sliceToSort)[right], (*sliceToSort)[pivot]

	for i, _ := range *sliceToSort {
		if (*sliceToSort)[i] < (*sliceToSort)[right] {
			(*sliceToSort)[left], (*sliceToSort)[i] = (*sliceToSort)[i], (*sliceToSort)[left]
			left++
		}
	}

	(*sliceToSort)[left], (*sliceToSort)[right] = (*sliceToSort)[right], (*sliceToSort)[left]

	leftSlice := (*sliceToSort)[:left]
	rightSlice := (*sliceToSort)[left+1:]

	quicksort(&leftSlice)
	quicksort(&rightSlice)
}
