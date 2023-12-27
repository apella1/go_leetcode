package easy

func RomanToInteger(roman string) int {
	value := 0
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"D": 50,
		"C": 100,
		"M": 1000,
	}
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanValues[string(roman[i])] < romanValues[string(roman[i+1])] {
			// todo: difference between string and rune and the default when iterating over a string
			value -= romanValues[string(roman[i])]
			continue
		}
		value += romanValues[string(roman[i])]
	}
	return value
}

/*
todo break, continue -> do multiple loop practice exercises
*/
