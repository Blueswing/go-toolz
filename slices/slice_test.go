package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestAll(t *testing.T) {
	res := All()
	assert.True(t, res)
	res = All(true)
	assert.True(t, res)
	res = All(true, false)
	assert.False(t, res)
}

func TestAny(t *testing.T) {
	res := Any()
	assert.False(t, res)
	res = Any(true)
	assert.True(t, res)
	res = Any(true, false)
	assert.True(t, res)
	res = Any(false, false)
	assert.False(t, res)
}

func TestChunk(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := Chunk(data, 3)
	assert.Equal(t, 3, len(res))
}

func TestConcat(t *testing.T) {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
	res := Concat(data)
	assert.Equal(t, 8, len(res))

}

func TestSum(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	intRes := Sum(ints)
	assert.Equal(t, 15, intRes)

	ptrs := []uintptr{1, 2, 3, 4, 5}
	ptrRes := Sum(ptrs)
	assert.Equal(t, uintptr(15), ptrRes)

	floats := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
	floatRes := Sum(floats)
	assert.InDelta(t, 6.5, floatRes, 1e-16)

	comps := []complex128{1 + 2i, 3 + 4i, 5 + 6i}
	compRes := Sum(comps)
	assert.Equal(t, 9+12i, compRes)
}

func TestProduct(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	intRes := Product(ints)
	assert.Equal(t, 120, intRes)

	comps := []complex128{1 + 2i, 3 + 4i, 5 + 6i}
	compRes := Product(comps)
	assert.Equal(t, -85+20i, compRes)
}

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ok := Contains(nums, 1)
	assert.True(t, ok, "expect true")
	ok = Contains(nums, 0)
	assert.False(t, ok, "expect false")

	strs := []string{"aaa", "bbb", "ccc"}
	ok = Contains(strs, "aaa")
	assert.True(t, ok, "expect true")
	ok = Contains(strs, "ddd")
	assert.False(t, ok, "expect false")
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	res := Reduce(s, func(x, y int) int { return x + y }, 0)
	assert.Equal(t, 15, res)
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	exp := []string{"1", "2", "3", "4", "5"}
	res := Map(s, func(i int) string { return strconv.Itoa(i) })
	ans := slices.EqualFunc(res, exp, func(x, y string) bool { return x == y })
	assert.True(t, ans)
}

func TestGroupBy(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := GroupBy(s, func(x int) int { return x % 2 })
	fmt.Println(res)
}

func TestPivot(t *testing.T) {
	s := []map[string]interface{}{
		{"k": 1, "v": "one"},
		{"k": 2, "v": "two"},
		{"k": 3, "v": "three"},
		{"k": 4, "v": "four"},
		{"k": 5, "v": "five"},
	}
	res := Pivot(s, func(x map[string]interface{}) (int, string) { return x["k"].(int), x["v"].(string) })
	fmt.Println(res)
}
