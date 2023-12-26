package main

import (
	"fmt"

	"github.com/apella1/go_leetcode/easy"
)

func main() {
	fmt.Println()
	fmt.Println("Easy")
	resultingSlice := easy.TwoSum([]int{2, 4, 5, 6}, 10)
	fmt.Println(resultingSlice)
	fmt.Println(easy.TwoSumLoops([]int{2, 4, 5, 6}, 10))
	fmt.Println(easy.RomanToInteger("IX"))
}
