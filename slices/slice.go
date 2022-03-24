package slices

import (
	"github.com/blueswing/go-toolz/defs"
	"golang.org/x/exp/constraints"
)

// Contains
func Contains[T any](s []T, elem T, eq defs.BinaryFunc[T, T, bool]) bool {
	for _, x := range s {
		if eq(x, elem) {
			return true
		}
	}
	return false
}

// ZAccumulate
func ZAccumulate[T any](initVal T, seq []T, acc defs.BinaryFunc[T, T, T]) T {
	retval := initVal
	for _, item := range seq {
		retval = acc(retval, item)
	}
	return retval
}

// Accumulate
func Accumulate[T any](initVal T, seq []T, acc defs.BinaryFunc[T, T, T]) T {
	retval := initVal
	for _, item := range seq {
		retval = acc(retval, item)
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
func GroupBy[T any](seq []T, f defs.UnaryFunc[T, any]) map[any][]T {
	retval := make(map[any][]T)
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
