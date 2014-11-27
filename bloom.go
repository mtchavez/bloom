package bloom

import (
	"hash"
	"hash/fnv"

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
