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
