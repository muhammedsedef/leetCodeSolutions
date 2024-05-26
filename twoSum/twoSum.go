package twoSum

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		refNum := target - num
		if index, exist := numMap[refNum]; exist {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{}
}

func RunTest() bool {
	for _, data := range GetData() {
		result := twoSum(data.Input, data.Target)
		if !reflect.DeepEqual(result, data.ExpectedResult) {
			fmt.Printf("Result: %v - ExpectedResult: %v\n", result, data.ExpectedResult)
			return false
		}
	}
	return true
}
