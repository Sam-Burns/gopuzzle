package gameMap

import "sort"

func CountLocations(ordinatesVisited *OrdinatesVisitedType) int {

	count := 0

	sort.Ints(ordinatesVisited[])

	for _, value := range ordinatesVisited {
		if value != 0 {
			count += 1
		}
	}

	return count
}

func SortOrdinates(ordinatesToSort *OrdinatesVisitedType) {



	for index, currentValue := range ordinatesToSort {

		int j = i - 1;

		/* Move elements of arr[0..i-1], that are
		   greater than key, to one position ahead
		   of their current position */
		while (j >= 0 && arr[j] > currentValue) {
			arr[j + 1] = arr[j];
			j = j - 1;
		}
		arr[j + 1] = key;
	}
}
