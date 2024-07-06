package romanToInt

type Data struct {
	Input          string
	ExpectedResult int
}

func GetData() []Data {
	return []Data{
		{Input: "III", ExpectedResult: 3},
		{Input: "LVIII", ExpectedResult: 58},
		{Input: "MCMXCIV", ExpectedResult: 1994},
	}
}
