package easy

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		word := ""
		if i%3 == 0 {
			word += "Fizz"
		}
		if i%5 == 0 {
			word += "Buzz"
		}
		if word == "" {
			fmt.Println(i)
		}
		fmt.Println(word)
	}
}
