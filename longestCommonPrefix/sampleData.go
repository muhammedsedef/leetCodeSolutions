package longestCommonPrefix

type Data struct {
	Input          []string
	ExpectedResult string
}

func GetData() []Data {
	return []Data{
		{Input: []string{"flower", "flow", "flight"}, ExpectedResult: "fl"},
		{Input: []string{"dog", "racecar", "car"}, ExpectedResult: ""},
	}
}
