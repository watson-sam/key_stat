// Package key_stats implements calculates "streak" length over a given number of float64 values.
package key_stat

import (
	"fmt"
	"math"
)

// StreakObject - the struct that holds the 'settings' and current streak length.
type StreakObject struct {
	value           int8
	ignoreNanValues bool
	ignoreInfValues bool
}

// SetIgnoreInfValues - controls if we want to ignore non number values when producing the outputs
// of any calculations
func (so *StreakObject) SetIgnoreNanValues(ignoreNanValues bool) {
	so.ignoreNanValues = ignoreNanValues
}

// SetIgnoreInfValues - controls if we want to ignore infinites (both positive and negative values)
// when producing the outputs of any calculations
func (so *StreakObject) SetIgnoreInfValues(ignoreInfValues bool) {
	so.ignoreInfValues = ignoreInfValues
}

// Add - if given value meets the given conditions, append to the values used in the calculation,
// adjusting this so it it relevant for the supplied window
func (so *StreakObject) Add(value float64) {
	if so.ignoreNanValues && math.IsNaN(value) {
		return
	}
	if so.ignoreInfValues && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return
	}

	if (value == 1.0) || (math.IsInf(value, 1)) {
		so.value++
	} else if (value == 0.0) || (math.IsInf(value, -1)) || math.IsNaN(value) {
		so.value = 0
	} else {
		panic("Supplied `value` argument is not valid - must be -inf, 0, 1, inf or nan, received value: " + fmt.Sprintf("%f", value))
	}
}
func (so *StreakObject) Value() int8 {
	return so.value
}

// NewStreakObject - set up a new rolling object with a supplied window with the default settings
func NewStreakObject() *StreakObject {
	return &StreakObject{
		ignoreNanValues: ignoreNanValuesDefault,
		ignoreInfValues: ignoreInfValuesDefault,
	}
}
