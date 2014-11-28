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

func Test_OptimalK_ZeroN(t *testing.T) {
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
