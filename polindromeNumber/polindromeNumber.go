package polindromeNumber

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	refX := x
	reversedX := 0

	for refX > 0 {
		reversedX = reversedX*10 + refX%10
		refX = refX / 10
	}

	return reversedX == x
}

func RunTest() bool {
	for _, data := range GetData() {
		result := isPalindrome(data.Input)
		if result != data.ExpectedResult {
			fmt.Printf("Result: %v - ExpectedResult: %v\n", result, data.ExpectedResult)
			return false
		}
	}
	return true
}
