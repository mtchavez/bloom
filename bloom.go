package bloom

import (
	"hash"
	"hash/fnv"
	"math"

	"github.com/mtchavez/bitset"
)

type Filter struct {
	k uint64
	m uint64
	f *bitset.Bitset
	h hash.Hash64
}

func NewFilter(k, m uint64) *Filter {
	return &Filter{
		k: k,
		m: m,
		f: bitset.New(m),
		h: fnv.New64(),
	}
}

func OptimalK(m, n uint64) uint64 {
	if n <= 0 {
		return 0
	}
	return uint64(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
}
