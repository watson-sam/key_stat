package key_stat

import "math"

const (
	// By default nan values are ignored (if false nan values are treated as 0)
	ignoreNanValuesDefault bool = true
	// By default infinite values (both positive and negative) are ignored
	// (if false positive inf values are treated as 1, negative inf values as 0)
	ignoreInfValuesDefault bool = true
)

var samples = [20]float64{
	1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
}

var suspectSamples = [20]float64{
	1, 1, math.NaN(), 1, 1, 0, 0, 0, math.NaN(), 0, 1, 1, 1, math.Inf(-1), 1, 1, math.NaN(), 1, 1, math.Inf(1),
}
