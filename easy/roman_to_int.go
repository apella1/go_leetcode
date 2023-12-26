package easy

func RomanToInteger(roman string) int {
	var value int
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"D": 50,
		"C": 100,
		"M": 1000,
	}
	// IV
	for i, v := range roman {
		if i+1 < len(roman) && (romanValues[string(v)] < romanValues[string(roman[i+1])]) {
			// todo: difference between string and rune and the default when iterating over a string
			value -= romanValues[string(v)]
		}
		value += romanValues[string(v)]
	}
	return value
}
