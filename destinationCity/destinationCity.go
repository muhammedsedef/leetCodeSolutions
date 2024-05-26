package destinationCity

import "fmt"

func DestCity(paths [][]string) string {
	existMap := make(map[string]bool)
	for _, path := range paths {
		existMap[path[0]] = true
		existMap[path[1]] = existMap[path[1]]
	}
	for key, value := range existMap {
		if !value {
			return key
		}
	}
	return ""
}

func RunTest() bool {
	for _, route := range getAllRoutes() {
		result := DestCity(route.Path)
		expectedResult := route.ExpectedResult
		if result != expectedResult {
			fmt.Printf("Target is not matched: %s\nYour result: %s", expectedResult, result)
			return false
		}
	}
	return true
}
