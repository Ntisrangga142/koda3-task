package utils

func Temperature(temp float64, origin string, goals string) float64 {
	var comparison = map[string]float64{
		"C":  5,
		"F":  9,
		"K":  5,
		"R":  9,
		"Re": 4,
	}

	var freezingPoint = map[string]float64{
		"C":  0,
		"F":  32,
		"K":  273.15,
		"R":  491.67,
		"Re": 0,
	}

	var result float64 = (temp-freezingPoint[origin])/comparison[origin]*comparison[goals] + freezingPoint[goals]
	return result
}
