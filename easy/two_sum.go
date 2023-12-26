package easy

/*
* LeetCode TwoSum Problem
* Assumptions
* only one valid solution
**/

func TwoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	mapStore := make(map[int]int)
	// []int{2, 3, 4, 8}, 12
	for index, value := range numbers {
		// diff := target - value
		// todo -> checking if a value exists on a map without the key
		_, ok := mapStore[index]
		if !ok {
			mapStore[index] = value
		}
	}
	return result
}

// asymptotic time is O(n^2)
func TwoSumLoops(numbers []int, target int) [2]int {
	result := [2]int{}
	for a := 0; a < len(numbers); a++ {
		for b := 1; b < len(numbers); b++ {
			if numbers[a]+numbers[b] == target {
				// todo how to manipulate static arrays
				result[0], result[1] = a, b
				break
			}
		}
	}
	return result
}
