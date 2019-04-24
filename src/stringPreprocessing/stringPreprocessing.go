package stringPreprocessing

func PreprocessString(originalString *string) []rune  {

	allRunes := []rune(*originalString)

	originalLength := len(*originalString)

	index := 1

	stoppingPoint := originalLength - 3

	for index <= stoppingPoint {

		if allRunes[index] == 'N' && allRunes[index+1] == 'S' && allRunes[index+2] == 'N' {
			allRunes[index+1] = 'X'
			allRunes[index+2] = 'X'
		}

		index++
	}

	return allRunes
}
