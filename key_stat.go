// Package key_stats implements calculates "key stats" over a number of float64 values within a given window
package key_stat

import (
	"fmt"
	"math"
)

// KeyStatsObject - the struct that holds the 'settings' and current values.
type KeyStatsObject struct {
	values    []int8
	minWindow int8
	maxWindow int8
	cutoff    float32

	ignoreNanValues bool
	ignoreInfValues bool
}

// SetIgnoreInfValues - controls if we want to ignore non number values when producing the outputs
// of any calculations
func (kso *KeyStatsObject) SetIgnoreNanValues(ignoreNanValues bool) {
	kso.ignoreNanValues = ignoreNanValues
}

// SetIgnoreInfValues - controls if we want to ignore infinites (both positive and negative values)
// when producing the outputs of any calculations
func (kso *KeyStatsObject) SetIgnoreInfValues(ignoreInfValues bool) {
	kso.ignoreInfValues = ignoreInfValues
}

// Add - if given value meets the given conditions, append to the values used in the calculation,
// adjusting this so it it relevant for the supplied windows
func (kso *KeyStatsObject) Add(value float64) {
	if kso.ignoreNanValues && math.IsNaN(value) {
		return
	}
	if kso.ignoreInfValues && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return
	}

	if len(kso.values) >= int(kso.maxWindow) {
		kso.values = kso.values[1:len(kso.values)]
	}
	if (value == 1.0) || (math.IsInf(value, 1)) {
		kso.values = append(kso.values, 1)
	} else if (value == 0.0) || (math.IsInf(value, -1)) || math.IsNaN(value) {
		kso.values = append(kso.values, 0)
	} else {
		panic("Supplied `value` argument is not valid - must be -inf, 0, 1, inf or nan, received value: " + fmt.Sprintf("%f", value))
	}
}

// KeyStat - return current key stat values and if they are relevant given the cutoff
func (kso *KeyStatsObject) KeyStat() (bool, int, int) {
	if len(kso.values) < int(kso.minWindow) {
		return false, 0, 0
	}

	values := make([]int8, len(kso.values))
	copy(values, kso.values)

	r := len(values)
	for i := 0; i < r; i++ {
		num := 0
		for _, j := range values {
			num += int(j)
		}
		denom := len(values)
		if float64(num)/float64(denom) >= float64(kso.cutoff) {
			return true, num, denom
		}
		values = values[1:]
		if len(values) < int(kso.minWindow) {
			return false, 0, 0
		}
	}

	panic("Error calculating current KeyStat values")
}

// NewKeyStatObject - set up a new key stat object with a supplied windows, cutoff and the default settings
func NewKeyStatObject(minWindow int8, maxWindow int8, cutoff float32) *KeyStatsObject {
	if minWindow > maxWindow {
		panic("`minWindow` argument must be less than `maxWindow`")
	}
	return &KeyStatsObject{
		minWindow:       minWindow,
		maxWindow:       maxWindow,
		cutoff:          cutoff,
		ignoreNanValues: ignoreNanValuesDefault,
		ignoreInfValues: ignoreInfValuesDefault,
	}
}
