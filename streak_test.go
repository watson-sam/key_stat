package key_stat

import (
	"testing"
)

func TestStreak(t *testing.T) {
	so := NewStreakObject()
	for _, f := range samples {
		so.Add(f)
	}
	result := so.Value()
	expected := int8(4)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestStreak2(t *testing.T) {
	so := NewStreakObject()
	for _, f := range suspectSamples {
		so.Add(f)
	}
	result := so.Value()
	expected := int8(7)
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
	result := so.Value()
	expected := int8(5)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}
