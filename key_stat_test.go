package key_stat

import (
	"testing"
)

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
