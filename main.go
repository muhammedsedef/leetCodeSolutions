package main

import (
	"fmt"
	"leetCodeSolutions/destinationCity"
	"leetCodeSolutions/longestCommonPrefix"
	"leetCodeSolutions/polindromeNumber"
	romanToInt "leetCodeSolutions/romanToInteger"
	"leetCodeSolutions/twoSum"
)

func main() {
	// destination city solution
	if destinationCityResult := destinationCity.RunTest(); destinationCityResult {
		fmt.Println("Destination city test successful")
	}

	//two sum solution
	if twoSumResult := twoSum.RunTest(); twoSumResult {
		fmt.Println("Two Sum test successful")
	}

	//polindrome number solution
	if polindromeNumberResult := polindromeNumber.RunTest(); polindromeNumberResult {
		fmt.Println("Palindrome number test successful")
	}

	//romanToInteger
	if romanToIntResult := romanToInt.RunTest(); romanToIntResult {
		fmt.Println("Roman To Integer test successful")
	}

	// longest common prefix
	if longestCommonPrefixResult := longestCommonPrefix.RunTest(); longestCommonPrefixResult {
		fmt.Println("Longest CommonPrefix test successful")
	}

}
