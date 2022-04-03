package slices

import (
	"github.com/yeefea/go-toolz/defs"
)

// All
func All(seq ...bool) bool {
	for _, x := range seq {
		if !x {
			return false
		}
	}
	return true
}

// AllFunc
func AllFunc[T any](pred defs.UnaryPred[T], seq ...T) bool {
	for _, x := range seq {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Any
func Any(seq ...bool) bool {
	for _, x := range seq {
		if x {
			return true
		}
	}
	return false
}

// AnyFunc
func AnyFunc[T any](pred defs.UnaryPred[T], seq ...T) bool {
	for _, x := range seq {
		if pred(x) {
			return true
		}
	}
	return false
}

// Sum returns the sum of given numbers
func Sum[T defs.Numeric](seq []T) T {
	var retval T
	for _, x := range seq {
		retval += x
	}
	return retval
}

// Product returns the product of given numbers
func Product[T defs.Numeric](seq []T) T {
	var retval T = 1
	for _, x := range seq {
		retval *= x
	}
	return retval
}

func Between[N defs.Ordered](x, sub, sup N) bool {
	return x >= sub && x <= sup
}

func Clamp[N defs.Ordered](x, sub, sup N) N {
	if x < sub {
		return sub
	}
	if x > sup {
		return sup
	}
	return x
}

// Accumulate
func Accumulate[T any](initVal T, seq []T, acc defs.BinaryFunc[T, T, T]) T {
	retval := initVal
	for _, item := range seq {
		retval = acc(retval, item)
	}
	return retval
}

// Min
func Min[T defs.Ordered](seq []T) T {
	ret := seq[0]
	for i := 1; i < len(seq); i++ {
		if seq[i] < ret {
			ret = seq[i]
		}
	}
	return ret
}

// MinFunc
func MinFunc[T any](seq []T, less defs.BinaryPred[T, T]) T {
	ret := seq[0]
	for i := 1; i < len(seq); i++ {
		if less(seq[i], ret) {
			ret = seq[i]
		}
	}
	return ret
}

// ArgMin
func ArgMin[T defs.Ordered](seq []T) int {
	idx := -1
	if len(seq) > 0 {
		min := seq[0]
		for i := 1; i < len(seq); i++ {
			if seq[i] < min {
				min = seq[i]
				idx = i
			}
		}
	}
	return idx
}

// ArgMinFunc
func ArgMinFunc[T defs.Ordered](seq []T, less defs.BinaryPred[T, T]) int {
	idx := -1
	if len(seq) > 0 {
		min := seq[0]
		for i := 1; i < len(seq); i++ {
			if less(seq[i], min) {
				min = seq[i]
				idx = i
			}
		}
	}
	return idx
}

// Max
func Max[T defs.Ordered](seq []T) T {
	ret := seq[0]
	for i := 1; i < len(seq); i++ {
		if seq[i] > ret {
			ret = seq[i]
		}
	}
	return ret
}

// MaxFunc
func MaxFunc[T any](seq []T, larger defs.BinaryPred[T, T]) T {
	ret := seq[0]
	for i := 1; i < len(seq); i++ {
		if larger(seq[i], ret) {
			ret = seq[i]
		}
	}
	return ret
}

// ArgMax
func ArgMax[T defs.Ordered](seq []T) int {
	idx := -1
	if len(seq) > 0 {
		max := seq[0]
		for i := 1; i < len(seq); i++ {
			if seq[i] > max {
				max = seq[i]
				idx = i
			}
		}
	}
	return idx
}

// ArgMaxFunc
func ArgMaxFunc[T defs.Ordered](seq []T, larger defs.BinaryPred[T, T]) int {
	idx := -1
	if len(seq) > 0 {
		min := seq[0]
		for i := 1; i < len(seq); i++ {
			if larger(seq[i], min) {
				min = seq[i]
				idx = i
			}
		}
	}
	return idx
}

// Contains
func Contains[T comparable](s []T, elem T) bool {
	for _, x := range s {
		if x == elem {
			return true
		}
	}
	return false
}

// ContainsFunc
func ContainsFunc[T any](s []T, elem T, eq defs.BinaryFunc[T, T, bool]) bool {
	for _, x := range s {
		if eq(x, elem) {
			return true
		}
	}
	return false
}

// Chunk
func Chunk[T any](s []T, size int) [][]T {
	retval := make([][]T, 0)
	for i := 0; i < len(s); i += size {
		retval = append(retval, s[i:i+size])
	}
	return retval
}

func Concat[T any](seqs ...[]T) []T {
	var c int
	for _, seq := range seqs {
		c += len(seq)
	}
	retval := make([]T, 0, c)
	for _, seq := range seqs {
		retval = append(retval, seq...)
	}
	return retval
}

// Pivot
func Pivot[T any, K comparable, O any](seq []T, f defs.PivotFunc[T, K, O]) map[K]O {
	retval := make(map[K]O)
	for _, x := range seq {
		key, out := f(x)
		retval[key] = out
	}
	return retval
}

// Unpivot
func Unpivot[K comparable, V any, O any](dict map[K]V, f defs.UnpivotFunc[K, V, O]) []O {
	retval := make([]O, 0, len(dict))
	for k, v := range dict {
		x := f(k, v)
		retval = append(retval, x)
	}
	return retval
}

// Map
func Map[I any, O any](seq []I, f defs.UnaryFunc[I, O]) []O {
	retval := make([]O, len(seq))
	for i, x := range seq {
		retval[i] = f(x)
	}
	return retval
}

// GroupBy
func GroupBy[T any, K comparable](seq []T, f defs.UnaryFunc[T, K]) map[K][]T {
	retval := make(map[K][]T)
	for _, x := range seq {
		key := f(x)
		retval[key] = append(retval[key], x)
	}
	return retval
}

// Join
func Join[T1 any, T2 any](seq1 []T1, seq2 []T2, eq defs.BinaryPred[T1, T2]) []defs.Pair[T1, T2] {
	retval := make([]defs.Pair[T1, T2], 0)
	for _, x := range seq1 {
		for _, y := range seq2 {
			if eq(x, y) {
				retval = append(retval, defs.Pair[T1, T2]{First: x, Second: y})
			}
		}
	}
	return retval
}

// FromPairs
func FromPairs[KeyType comparable, ValueType any](pairs []defs.Pair[KeyType, ValueType]) map[KeyType]ValueType {
	retval := make(map[KeyType]ValueType)
	for _, p := range pairs {
		retval[p.First] = p.Second
	}
	return retval
}

// ToPairs
func ToPairs[KeyType comparable, ValueType any](dict map[KeyType]ValueType) []defs.Pair[KeyType, ValueType] {
	retval := make([]defs.Pair[KeyType, ValueType], 0, len(dict))
	for k, v := range dict {
		retval = append(retval, defs.Pair[KeyType, ValueType]{First: k, Second: v})
	}
	return retval
}
