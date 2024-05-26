package twoSum

type Data struct {
	Input          []int
	Target         int
	ExpectedResult []int
}

func GetData() []Data {
	return []Data{
		{Input: []int{2, 7, 11, 15}, Target: 9, ExpectedResult: []int{0, 1}},
		{Input: []int{3, 2, 4}, Target: 6, ExpectedResult: []int{1, 2}},
		{Input: []int{3, 3}, Target: 6, ExpectedResult: []int{0, 1}},
	}
}
