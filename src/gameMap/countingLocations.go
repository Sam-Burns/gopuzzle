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

func quicksort(a *[]uint64) {
	if len(*a) < 2 {
		return// *a
	}

	left, right := 0, len(*a)-1

	pivot := rand.Int() % len(*a)

	(*a)[pivot], (*a)[right] = (*a)[right], (*a)[pivot]

	for i, _ := range *a {
		if (*a)[i] < (*a)[right] {
			(*a)[left], (*a)[i] = (*a)[i], (*a)[left]
			left++
		}
	}

	(*a)[left], (*a)[right] = (*a)[right], (*a)[left]

	leftSlice := (*a)[:left]
	rightSlice := (*a)[left+1:]

	quicksort(&leftSlice)
	quicksort(&rightSlice)
}
