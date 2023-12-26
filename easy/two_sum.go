package easy

import "fmt"

/*
* LeetCode TwoSum Problem
* Assumptions
* only one valid solution
**/

func TwoSum(numbers []int, target int) ([]int, map[int]int) {
	result := make([]int, len(numbers))
	mapStore := make(map[int]int)
	for index, value := range numbers {
		// check if value is in mapStore
		// if true return the index of the value in the map store and the current index
		// otherwise store the current value in the map store and continue with the loop
		diff := target - value
		fmt.Println(diff)
		mapStore[index] = value
	}
	return result, mapStore
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
