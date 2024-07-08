package longestCommonPrefix

import (
	"fmt"
	"reflect"
)

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < length; i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" {
			break
		}
	}

	return prefix

}

func commonPrefix(str1, str2 string) string {
	minLength := len(str1)
	if len(str2) < minLength {
		minLength = len(str2)
	}

	for i := 0; i < minLength; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}

	return str1[:minLength]
}

func RunTest() bool {
	for _, data := range GetData() {
		result := longestCommonPrefix(data.Input)
		if !reflect.DeepEqual(result, data.ExpectedResult) {
			fmt.Printf("Result: %s - ExpectedResult: %s\n", result, data.ExpectedResult)
			return false
		}
	}
	return true
}
