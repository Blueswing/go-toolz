package slices

import (
	"github.com/yeefea/go-toolz/defs"
)

// All returns true if all elements in seq are true or if seq is empty
func All(seq ...bool) bool {
	for _, x := range seq {
		if !x {
			return false
		}
	}
	return true
}

// AllPredicates returns true if all predicates in preds true or if preds is empty
func AllPredicates[T any](val T, preds ...defs.UnaryPred[T]) bool {
	for _, f := range preds {
		if !f(val) {
			return false
		}
	}
	return true
}

// Any returns false if all elements in seq are false or if seq is empty
func Any(seq ...bool) bool {
	for _, x := range seq {
		if x {
			return true
		}
	}
	return false
}

// AnyFunc returns false if all predicates in preds returns false or if preds is empty
func AnyPredicates[T any](val T, preds ...defs.UnaryPred[T]) bool {
	for _, f := range preds {
		if f(val) {
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

// Chunk
func Chunk[T any](s []T, size int) [][]T {
	retval := make([][]T, 0)
	for start := 0; start < len(s); start += size {
		end := start + size
		if end > len(s) {
			end = len(s)
		}
		retval = append(retval, s[start:end])
	}
	return retval
}

// Concat
func Concat[T any](seqs [][]T) []T {
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

// Contains if seq contains elem
func Contains[T comparable](seq []T, elem T) bool {
	for _, x := range seq {
		if x == elem {
			return true
		}
	}
	return false
}

// ContainsFunc
func ContainsFunc[T any](seq []T, pred defs.UnaryPred[T]) bool {
	for _, x := range seq {
		if pred(x) {
			return true
		}
	}
	return false
}

// Count
func Count[T comparable](seq []T, val T) int {
	cnt := 0
	for _, x := range seq {
		if x == val {
			cnt++
		}
	}
	return cnt
}

// CountFunc
func CountFunc[T any](seq []T, pred defs.UnaryPred[T]) int {
	cnt := 0
	for _, x := range seq {
		if pred(x) {
			cnt++
		}
	}
	return cnt
}

// Counter
func Counter[T comparable](seq []T) map[T]int {
	retval := make(map[T]int)
	for _, x := range seq {
		retval[x] += 1
	}
	return retval
}

// NewSet
func NewSet[T comparable](seq []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, x := range seq {
		set[x] = struct{}{}
	}
	return set
}

// Difference return the difference between two sequences
func Difference[T comparable](seq1, seq2 []T) ([]T, []T) {
	set1 := NewSet(seq1)
	set2 := NewSet(seq2)
	diff1 := make([]T, 0)
	diff2 := make([]T, 0)
	for _, x := range seq1 {
		if _, ok := set2[x]; !ok {
			diff1 = append(diff1, x)
		}
	}
	for _, x := range seq2 {
		if _, ok := set1[x]; !ok {
			diff2 = append(diff2, x)
		}
	}
	return diff1, diff2
}

// Equal
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc
func EqualFunc[T1 any, T2 any](s1 []T1, s2 []T2, eq defs.BinaryPred[T1, T2]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Filter
func Filter[T any](seq []T, pred defs.UnaryPred[T]) []T {
	retval := make([]T, 0)
	for _, x := range seq {
		if pred(x) {
			retval = append(retval, x)
		}
	}
	return retval
}

// Find
func Find[T any](seq []T, pred defs.UnaryPred[T]) (T, bool) {
	for _, x := range seq {
		if pred(x) {
			return x, true
		}
	}
	var retval T
	return retval, false
}

// Index
func Index[T comparable](seq []T, val T) int {
	idx := -1
	for i, x := range seq {
		if x == val {
			idx = i
			break
		}
	}
	return idx
}

// IndexFunc
func IndexFunc[T any](seq []T, pred defs.UnaryPred[T]) int {
	idx := -1
	for i, x := range seq {
		if pred(x) {
			idx = i
			break
		}
	}
	return idx
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

// Reduce
func Reduce[T any](seq []T, acc defs.ReduceFunc[T], initVal T) T {
	retval := initVal
	for _, item := range seq {
		retval = acc(retval, item)
	}
	return retval
}

// Reverse reverse seq in place
func Reverse[T any](seq []T) {
	for left, right := 0, len(seq)-1; left < right; left, right = left+1, right-1 {
		seq[left], seq[right] = seq[right], seq[left]
	}
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

// Zip
func Zip[T1 any, T2 any](seq1 []T1, seq2 []T2) []defs.Pair[T1, T2] {
	var commonLen int
	if len(seq1) < len(seq2) {
		commonLen = len(seq1)
	} else {
		commonLen = len(seq2)
	}
	retval := make([]defs.Pair[T1, T2], 0, commonLen)
	for i := 0; i < commonLen; i++ {
		retval = append(retval, defs.Pair[T1, T2]{First: seq1[i], Second: seq2[i]})
	}
	return retval
}

// Unzip
func Unzip[T1 any, T2 any](seq []defs.Pair[T1, T2]) ([]T1, []T2) {
	ret1 := make([]T1, 0, len(seq))
	ret2 := make([]T2, 0, len(seq))
	for _, x := range seq {
		ret1 = append(ret1, x.First)
		ret2 = append(ret2, x.Second)
	}
	return ret1, ret2
}
