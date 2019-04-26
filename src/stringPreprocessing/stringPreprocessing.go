package stringPreprocessing

import "fmt"

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

	{search: []rune{'E', 'S', 'N'}, replacement: []rune{'3', 'N', 'X'}, removals: 1},
	{search: []rune{'E', 'N', 'S'}, replacement: []rune{'9', 'S', 'X'}, removals: 1},
	{search: []rune{'W', 'S', 'N'}, replacement: []rune{'1', 'N', 'X'}, removals: 1},
	{search: []rune{'W', 'N', 'S'}, replacement: []rune{'7', 'S', 'X'}, removals: 1},
	{search: []rune{'S', 'E', 'W'}, replacement: []rune{'3', 'W', 'X'}, removals: 1},
	{search: []rune{'S', 'W', 'E'}, replacement: []rune{'1', 'E', 'X'}, removals: 1},
	{search: []rune{'N', 'E', 'W'}, replacement: []rune{'9', 'W', 'X'}, removals: 1},
	{search: []rune{'N', 'W', 'E'}, replacement: []rune{'7', 'E', 'X'}, removals: 1},

	{search: []rune{'S', 'N', 'W'}, replacement: []rune{'S', '7', 'X'}, removals: 1},
	{search: []rune{'S', 'N', 'E'}, replacement: []rune{'S', '9', 'X'}, removals: 1},
	{search: []rune{'N', 'S', 'E'}, replacement: []rune{'N', '3', 'X'}, removals: 1},
	{search: []rune{'N', 'S', 'W'}, replacement: []rune{'N', '1', 'X'}, removals: 1},
	{search: []rune{'E', 'W', 'N'}, replacement: []rune{'E', '7', 'X'}, removals: 1},
	{search: []rune{'E', 'W', 'S'}, replacement: []rune{'E', '1', 'X'}, removals: 1},
	{search: []rune{'W', 'E', 'N'}, replacement: []rune{'W', '9', 'X'}, removals: 1},
	{search: []rune{'W', 'E', 'S'}, replacement: []rune{'W', '3', 'X'}, removals: 1},

	{search: []rune{'N', 'N', 'S'}, replacement: []rune{'8', 'S', 'X'}, removals: 1},
	{search: []rune{'S', 'S', 'N'}, replacement: []rune{'2', 'N', 'X'}, removals: 1},
	{search: []rune{'E', 'E', 'W'}, replacement: []rune{'6', 'W', 'X'}, removals: 1},
	{search: []rune{'W', 'W', 'E'}, replacement: []rune{'4', 'E', 'X'}, removals: 1},

	{search: []rune{'S', 'N', 'N'}, replacement: []rune{'S', '8', 'X'}, removals: 1},
	{search: []rune{'N', 'S', 'S'}, replacement: []rune{'N', '2', 'X'}, removals: 1},
	{search: []rune{'E', 'W', 'W'}, replacement: []rune{'E', '4', 'X'}, removals: 1},
	{search: []rune{'W', 'E', 'E'}, replacement: []rune{'W', '6', 'X'}, removals: 1},
}

func PreprocessString(originalString *string) (allRunes []rune, directionsNullified int) {
	return fastPreprocessString(originalString)
}

func slowPreprocessString(originalString *string) (allRunes []rune, directionsNullified int) {

	allRunes = []rune(*originalString)

	originalLength := len(*originalString)

	index := 1

	stoppingPoint := originalLength - 3

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

	if len(allRunes) != 0 {
		length := float32(len(allRunes))
		delta := float32(directionsNullified)
		compressionRatio := delta / length
		fmt.Printf("%.2f \n", compressionRatio*100)
	} else {
		fmt.Println("no op")
	}

	return allRunes, directionsNullified
}

func fastPreprocessString(originalString *string) (allRunes []rune, directionsNullified int) {

	allRunes = []rune(*originalString)

	originalLength := len(*originalString)

	index := 1

	stoppingPoint := originalLength - 3

	directionsNullified = 0

	for index <= stoppingPoint {

		if allRunes[index] == 'N' {
			if allRunes[index+1] == 'S' {
				if allRunes[index+2] == 'N' {
					allRunes[index+1] = 'X'
					allRunes[index+2] = 'X'
					directionsNullified += 2
				} else if allRunes[index+2] == 'E' {
					allRunes[index+1]   = '3'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'W' {
					allRunes[index+1]   = '1'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'S' {
					allRunes[index+1]   = '2'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'E' {
				if allRunes[index+2] == 'W' {
					allRunes[index]   = '9'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'W' {
				if allRunes[index+2] == 'E' {
					allRunes[index]   = '7'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'N' {
				if allRunes[index+2] == 'S' {
					allRunes[index]   = '8'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			}

		} else if allRunes[index] == 'S' {
			if allRunes[index+1] == 'N' {
				if allRunes[index+2] == 'W' {
					allRunes[index+1] = 'X'
					allRunes[index+2] = 'X'
					directionsNullified += 2
				} else if allRunes[index+2] == 'S' {
					allRunes[index+1] = 'X'
					allRunes[index+2] = 'X'
					directionsNullified += 2
				} else if allRunes[index+2] == 'E' {
					allRunes[index+1]   = '9'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'N' {
					allRunes[index+1]   = '8'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'E' {
				if allRunes[index+2] == 'W' {
					allRunes[index]   = '3'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'W' {
				if allRunes[index+2] == 'E' {
					allRunes[index]   = '1'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'S' {
				if allRunes[index+2] == 'N' {
					allRunes[index]   = '2'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			}

		} else if allRunes[index] == 'E' {
			if allRunes[index+1] == 'W' {
				if allRunes[index+2] == 'E' {
					allRunes[index+1] = 'X'
					allRunes[index+2] = 'X'
					directionsNullified += 2
				} else if allRunes[index+2] == 'N' {
					allRunes[index+1]   = '7'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'S' {
					allRunes[index+1] = '1'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'W' {
					allRunes[index+1] = '4'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'S' {
				if allRunes[index+2] == 'N' {
					allRunes[index]   = '3'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'N' {
				if allRunes[index+2] == 'S' {
					allRunes[index]   = '9'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'E' {
				if allRunes[index+2] == 'W' {
					allRunes[index]   = '6'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			}

		} else if allRunes[index] == 'W' {
			if allRunes[index+1] == 'E' {
				if allRunes[index+2] == 'W' {
					allRunes[index+1] = 'X'
					allRunes[index+2] = 'X'
					directionsNullified += 2
				} else if allRunes[index+2] == 'N' {
					allRunes[index+1] = '9'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'S' {
					allRunes[index+1] = '3'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				} else if allRunes[index+2] == 'E' {
					allRunes[index+1] = '6'
					allRunes[index+2] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'S' {
				if allRunes[index+2] == 'N' {
					allRunes[index]   = '1'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'N' {
				if allRunes[index+2] == 'S' {
					allRunes[index]   = '7'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			} else if allRunes[index+1] == 'W' {
				if allRunes[index+2] == 'E' {
					allRunes[index]   = '4'
					allRunes[index+1] = 'X'
					directionsNullified += 1
				}
			}
		}

		index++
	}

	return

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
