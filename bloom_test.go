package bloom

import (
	"testing"
)

func Test_NewFilter(t *testing.T) {
	var k uint64 = 10
	var m uint64 = 100
	f := NewFilter(k, m)
	if f.k != k {
		t.Errorf("Expected k to be set to %+v but got %+v", k, f.k)
	}
	if f.m != m {
		t.Errorf("Expected m to be set to %+v but got %+v", m, f.m)
	}
}

func Test_OptimalK_zeroN(t *testing.T) {
	k := OptimalK(1000, 0)
	if k != 0 {
		t.Errorf("Expected k to be zero if n is <= 0")
	}
}

func Test_OptimalK(t *testing.T) {
	k := OptimalK(1e9, 1e6)
	if k != 694 {
		t.Errorf("Expected k to be 694 but got %+v", k)
	}

	k = OptimalK(1e12, 1e11)
	if k != 7 {
		t.Errorf("Expected k to be 7 but got %+v", k)
	}
}

func Test_FalsePositiveProb_zeroN(t *testing.T) {
	prob := FalsePositiveProb(64, 1e9)
	if prob != 0.9999999692510076 {
		t.Errorf("Expected prob to be 0.99 but got %+v", prob)
	}
}
func Test_MForFalsePositiveProb(t *testing.T) {
	m := MForFalsePositiveProb(1e12, 0.000009)
	var expected uint64 = 24181940052532
	if m != expected {
		t.Errorf("Expected m to be but got %+v", m)
	}
}
