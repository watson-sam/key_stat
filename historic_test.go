package key_stat

import (
	"testing"
)

var histSamples = [10]float64{
	4, 2, 3, 3, 4, 7, 8, 10, 2, 3,
}

func TestHistoric(t *testing.T) {
	ho := NewHistoricObject("max")
	for _, f := range histSamples {
		ho.Add(f)
	}
	result := ho.Compare(9)
	expected := int8(0)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestHistoric2(t *testing.T) {
	ho := NewHistoricObject("max")
	for _, f := range histSamples {
		ho.Add(f)
	}
	result := ho.Compare(10)
	expected := int8(1)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestHistoric3(t *testing.T) {
	ho := NewHistoricObject("max")
	for _, f := range histSamples {
		ho.Add(f)
	}
	result := ho.Compare(11)
	expected := int8(2)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestHistoric4(t *testing.T) {
	ho := NewHistoricObject("min")
	for _, f := range histSamples {
		ho.Add(f)
	}
	result := ho.Compare(1)
	expected := int8(2)
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}
