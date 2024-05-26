package polindromeNumber

type Data struct {
	Input          int
	ExpectedResult bool
}

func GetData() []Data {
	return []Data{
		{Input: 121, ExpectedResult: true},
		{Input: -121, ExpectedResult: false},
		{Input: 10, ExpectedResult: false},
		{Input: 0, ExpectedResult: true},
	}
}
