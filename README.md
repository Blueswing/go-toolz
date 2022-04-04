# go-toolz
`go-toolz` is a functional library written in Go.

This library is inspired by [PyToolz](https://github.com/pytoolz/toolz) and it is implemented using the generic features of Go 1.18 and above.

## Requirements

Go 1.18+

## Installation

```bash
go get github.com/yeefea/go-toolz
```

## Usage

```go
import "github.com/yeefea/go-toolz"
```

## Examples

### All
```go
res := toolz.All(true, true, false, true)  // false
res := toolz.All(true, true, true)  // true
res := toolz.All() // true
```

### Any
```go
res := toolz.Any(true, false, false)  // true
res := toolz.Any(false, false)  // false
res := toolz.Any() // false
```

### Sum
```go
res := toolz.Sum([]int{1, 2, 3, 4, 5})  // 15
res := toolz.Sum([]float64{1.1, 1.2, 1.3, 1.4, 1.5})  // 6.5
res := toolz.Sum([]complex128{1 + 2i, 3 + 4i, 5 + 6i})  // 9 + 12i
```

### Product
```go
res := toolz.Product([]int{1, 2, 3, 4, 5})  // 120
res := toolz.Product([]complex128{1 + 2i, 3 + 4i, 5 + 6i})  // -85 + 20i
```

### Reduce
```go
data := []int{1, 2, 3}
res := slices.Reduce(data, func(x, y int) int { return 2*x + y }, 0)  // 11
```

### Map
```go
res := toolz.Map([]int{1, 2, 3}, func(x int) float64 { return float64(x * x) })
// [1.0 4.0 9.0 16.0 25.0]
```

### Min and max
```go
res := toolz.Min([]int{1, 3, 2, 4})  // 1
res := toolz.Max([]int{1, 3, 2, 4})  // 4
```

### Chunk
```go
res := toolz.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)
// [[1 2 3] [4 5 6] [7 8]]

res := toolz.Chunk([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, 4)
// [["1" "2" "3" "4"] ["5" "6" "7" "8"]]
```

### Concat
```go
data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
res := Concat(data)
// [1, 2, 3, 4, 5, 6, 7, 8]
```

### Contains
```go
res := Contains([]int{1, 2, 3, 4, 5}, 3)  // true
res := Contains([]int{1, 2, 3, 4, 5}, 0)  // false

strs := []string{"aaa", "bbb", "ccc"}
res := Contains(strs, "aaa")  // true
res := Contains(strs, "ddd")  // false
```

### GroupBy
```go
seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
res := toolz.GroupBy(seq, func(x int) int { return x % 2 })  
// map[0:[2 4 6 8] 1:[1 3 5 7]]
```

### Pivot
```go
s := []map[string]interface{}{
	{"k": 1, "v": "one"},
	{"k": 2, "v": "two"},
	{"k": 3, "v": "three"},
	{"k": 4, "v": "four"},
	{"k": 5, "v": "five"},
}
res := toolz.Pivot(s, func(x map[string]interface{}) (int, string) { return x["k"].(int), x["v"].(string) })  
// map[1:one 2:two 3:three 4:four 5:five]
```

### Zip
```go
data1 := []int{1, 2, 3}
data2 := []string{"one", "two", "three"}
res := toolz.Zip(data1, data2)
// [{1 one} {2 two} {3 three}]
```

### Unzip
```go
data := []defs.Pair[int, string]{
	{First: 1, Second: "one"},
	{First: 2, Second: "two"},
	{First: 3, Second: "three"},
}
res1, res2 := slices.Unzip(data)
// [1 2 3] [one two three]
```