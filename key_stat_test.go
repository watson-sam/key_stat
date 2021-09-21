package key_stat

import (
	"math"
	"testing"
)

var samples = [20]float64{
	1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
}

var suspectSamples = [20]float64{
	1, 1, math.NaN(), 1, 1, 0, 0, 0, math.NaN(), 0, 1, 1, 1, math.Inf(-1), 1, 1, math.NaN(), 1, 1, math.Inf(1),
}

func TestStreak(t *testing.T) {
	so := NewStreakObject()
	for _, f := range samples {
		so.Add(f)
	}
	result := so.Streak()
	expected := 4.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestStreak2(t *testing.T) {
	so := NewStreakObject()
	for _, f := range suspectSamples {
		so.Add(f)
	}
	result := so.Streak()
	expected := 7.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestStreak3(t *testing.T) {
	so := NewStreakObject()
	so.SetIgnoreInfValues(false)
	for _, f := range suspectSamples {
		so.Add(f)
	}
	result := so.Streak()
	expected := 5.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestKeyStat(t *testing.T) {
	so := NewKeyStatObject(5, 8, 0.5)
	for _, f := range samples {
		so.Add(f)
	}
	b, num, denom := so.KeyStat()
	numExp, denomExp := 4, 8
	if !b || (num != numExp) || (denom != denomExp) {
		t.Errorf("e.KeyStat() returns %v, %v, %v wanted %v, %v, %v", b, num, denom, true, numExp, denomExp)
	}
}

func TestKeyStat2(t *testing.T) {
	so := NewKeyStatObject(5, 12, 0.7)
	for _, f := range samples {
		so.Add(f)
	}
	b, num, denom := so.KeyStat()
	numExp, denomExp := 4, 5
	if !b || (num != numExp) || (denom != denomExp) {
		t.Errorf("e.KeyStat() returns %v, %v, %v wanted %v, %v, %v", b, num, denom, true, numExp, denomExp)
	}
}

func TestKeyStat3(t *testing.T) {
	so := NewKeyStatObject(6, 10, 0.5)
	so.SetIgnoreInfValues(false)
	for _, f := range suspectSamples {
		so.Add(f)
	}
	b, num, denom := so.KeyStat()
	numExp, denomExp := 8, 10
	if !b || (num != numExp) || (denom != denomExp) {
		t.Errorf("e.KeyStat() returns %v, %v, %v wanted %v, %v, %v", b, num, denom, true, numExp, denomExp)
	}
}
