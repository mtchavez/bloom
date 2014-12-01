package bloom

import (
	"hash"
	"hash/fnv"
	"math"

	"github.com/mtchavez/bitset"
)

type Filter struct {
	k     uint64
	m     uint64
	bset  *bitset.Bitset
	hashA hash.Hash64
	hashB hash.Hash64
	total uint64
}

func NewFilter(k, m uint64) *Filter {
	return &Filter{
		k:     k,
		m:     m,
		bset:  bitset.New(m),
		hashA: fnv.New64(),
		hashB: fnv.New64(),
	}
}

func OptimalK(m, n uint64) uint64 {
	if n <= 0 {
		return 0
	}
	return uint64(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
}

func FalsePositiveProb(m, n uint64) float64 {
	if n <= 0 {
		return 0
	}
	return math.Pow(math.E, -(float64(m) / (float64(n)) * math.Pow(math.Log(2), 2)))
}

func MForFalsePositiveProb(n uint64, p float64) uint64 {
	return uint64(math.Ceil(-((float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2))))
}

func (f *Filter) Add(data []byte) {
	for _, bit := range f.hashBits(data)[:f.k] {
		f.bset.Set(bit)
	}
	f.total++
}

func (f *Filter) Check(data []byte) bool {
	for _, bit := range f.hashBits(data)[:f.k] {
		if !f.bset.Test(bit) {
			return false
		}
	}
	return true
}

func (f *Filter) hashBits(data []byte) []uint64 {
	f.hashA.Reset()
	f.hashA.Write(data)
	a := f.hashA.Sum64()

	f.hashB.Reset()
	f.hashB.Write(data)
	b := f.hashB.Sum64()

	bits := make([]uint64, f.k)
	for i := uint64(0); i < f.k; i++ {
		bits[i] = (a + b*i) % f.m
	}
	return bits
}
