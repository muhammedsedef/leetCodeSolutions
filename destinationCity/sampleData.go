package destinationCity

type Route struct {
	Path           [][]string
	ExpectedResult string
}

func getAllRoutes() []Route {
	return []Route{
		{
			Path: [][]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			ExpectedResult: "Sao Paulo",
		},
		{
			Path: [][]string{
				{"B", "C"},
				{"D", "B"},
				{"C", "A"},
			},
			ExpectedResult: "A",
		},
		{
			Path: [][]string{
				{"A", "Z"},
			},
			ExpectedResult: "Z",
		},
		{
			Path: [][]string{
				{"pYyNGfBYbm", "wxAscRuzOl"},
				{"kzwEQHfwce", "pYyNGfBYbm"},
			},
			ExpectedResult: "wxAscRuzOl",
		},
	}
}
