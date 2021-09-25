package key_stat

import (
	"math"
)

// HistoricObject - the struct that holds the 'settings' and current values.
type HistoricObject struct {
	value float64
	//date     time.Time
	statType string
}

// Add a given value to the historic object if it is the latest version the is relevant for the stat type
func (ho *HistoricObject) Add(value float64) {
	if math.IsNaN(value) && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return
	}
	if (ho.statType == "max") && ((value >= ho.value) || (ho.value == -1000)) {
		ho.value = value
		//ho.date = date
	} else if (ho.statType == "min") && ((value <= ho.value) || (ho.value == 1000)) {
		ho.value = value
		//ho.date = date
	}
}

// Compare a given value with the current max or min stored in the historic object
func (ho *HistoricObject) Compare(value float64) int8 {
	if math.IsNaN(value) && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return 0
	}
	if value == ho.value {
		return 1
	} else if (ho.statType == "max") && (value > ho.value) {
		return 2
	} else if (ho.statType == "min") && (value < ho.value) {
		return 2
	}
	return 0
}

// NewHistoricObject - set up a new historic object to check for a given stat type
func NewHistoricObject(statType string) *HistoricObject {
	if statType == "max" {
		return &HistoricObject{
			statType: statType,
			value:    -1000,
		}
	} else if statType == "min" {
		return &HistoricObject{
			statType: statType,
			value:    1000,
		}
	}
	return &HistoricObject{}
}
