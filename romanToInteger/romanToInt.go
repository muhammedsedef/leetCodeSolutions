package romanToInt

import (
	"fmt"
	"reflect"
)

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	l := len(s)
	result := 0

	for i := 0; i < l; i++ {
		t := ""
		if i != l-1 {
			t = string(s[i]) + string(s[i+1])
			if val, exist := romanMap[t]; exist {
				result += val
				i++
				continue
			}
		}

		t = string(s[i])

		if val, exist := romanMap[t]; exist {
			result += val
		}

	}

	return result
}

func RunTest() bool {
	for _, data := range GetData() {
		result := romanToInt(data.Input)
		if !reflect.DeepEqual(result, data.ExpectedResult) {
			fmt.Printf("Result: %d - ExpectedResult: %d\n", result, data.ExpectedResult)
			return false
		}
	}
	return true
}
