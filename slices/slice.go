package slices

import (
	"github.com/blueswing/go-toolz/defs"
	"golang.org/x/exp/constraints"
)

// Sum
func Sum[T defs.Numeric](initVal T, seq []T) T {
	retval := initVal
	for _, x := range seq {
		retval += x
	}
	return retval
}

// Product
func Product[T defs.Numeric](initVal T, seq []T) T {
	retval := initVal
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

// Contains
func Contains[T any](s []T, elem T, eq defs.BinaryFunc[T, T, bool]) bool {
	for _, x := range s {
		if eq(x, elem) {
			return true
		}
	}
	return false
}

func Chunk[T any](s []T, size int) [][]T {
	retval := make([][]T, 0)
	for i := 0; i < len(s); i += size {
		retval = append(retval, s[i:i+size])
	}
	return retval
}

func Concat[T any](seqs ...[]T) []T {
	retval := make([]T, 0)
	for _, seq := range seqs {
		retval = append(retval, seq...)
	}
	return retval
}

// Pivot
func Pivot[T any, K comparable, O any](seq []T, f func(T) (K, O)) map[K]O {
	retval := make(map[K]O)
	for _, x := range seq {
		key, out := f(x)
		retval[key] = out
	}
	return retval
}

// Unpivot
func Unpivot[K comparable, V any, O any](dict map[K]V, f func(K, V) O) []O {
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

func ZClamp[T constraints.Ordered](max, min, val T) T {
	if val > max {
		return max
	}
	if val < min {
		return min
	}
	return val
}

func Join[T1 any, T2 any](seq1 []T1, seq2 []T2, eq defs.BinaryPredicate[T1, T2]) []defs.Pair[T1, T2] {
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
