package slices

import (
	"strconv"
	"testing"

	"github.com/blueswing/go-toolz/defs"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	ok := Contains(nums, 1, defs.ZEqual[int])
	assert.True(t, ok, "expect true")
	ok = Contains(nums, 0, defs.ZEqual[int])
	assert.False(t, ok, "expect false")

	strs := []string{"aaa", "bbb", "ccc"}
	ok = Contains(strs, "aaa", defs.ZEqual[string])
	assert.True(t, ok, "expect true")
	ok = Contains(strs, "ddd", defs.ZEqual[string])
	assert.False(t, ok, "expect false")
}

func TestAccumulate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	res := Accumulate(0, s, func(x, y int) int { return x + y })
	assert.Equal(t, 15, res)
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	exp := []string{"1", "2", "3", "4", "5"}
	res := Map[int, string](s, func(i int) string { return strconv.Itoa(i) })
	ans := slices.EqualFunc(res, exp, func(x, y string) bool { return x == y })
	assert.True(t, ans)
}
