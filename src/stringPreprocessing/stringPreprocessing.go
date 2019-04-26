package stringPreprocessing

type strippingPair struct {
	search []rune
	replacement []rune
	removals int
}

var strippingPairs = [...]strippingPair{
	{search: []rune{'N', 'S', 'N'}, replacement: []rune{'N', 'X', 'X'}, removals: 2},
	{search: []rune{'S', 'N', 'S'}, replacement: []rune{'S', 'X', 'X'}, removals: 2},
	{search: []rune{'E', 'W', 'E'}, replacement: []rune{'E', 'X', 'X'}, removals: 2},
	{search: []rune{'W', 'E', 'W'}, replacement: []rune{'W', 'X', 'X'}, removals: 2},
}

func PreprocessString(originalString *string) (allRunes []rune, directionsNullified int) {

	allRunes = []rune(*originalString)

	originalLength := len(*originalString)

	index := 1

	stoppingPoint := originalLength - 1

	directionsNullified = 0

	for index <= stoppingPoint {

		for _, aStrippingPair := range strippingPairs {
			if currentRuneIsStartOfStrippingPattern(&(aStrippingPair.search), &allRunes, &index, &originalLength) {

				for indexOffset, replacementRune := range aStrippingPair.replacement {
					allRunes[index + indexOffset] = replacementRune
				}

				directionsNullified += aStrippingPair.removals
			}
		}

		index++
	}

	return allRunes, directionsNullified
}

func currentRuneIsStartOfStrippingPattern(strippingPattern *[]rune, allRunes *[]rune, currentIndex *int, maxLength *int) bool {
	if *currentIndex + len(*strippingPattern) > *maxLength {
		return false
	}

	for strippingPatternIndex, strippingPatternValue := range *strippingPattern {
		if (*allRunes)[*currentIndex + strippingPatternIndex] != strippingPatternValue {
			return false
		}
	}
	return true
}
