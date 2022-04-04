# go-toolz
A functional library written in Go.

## Usage

```bash
go get github.com/yeefea/go-toolz
```

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
res := toolz.Sum([]complex128{1 + 2i, 3 + 4i, 5 + 6i})  // 9+12i
```

### Product
```go
res := toolz.Product([]int{1, 2, 3, 4, 5})  // 120
res := toolz.Product([]complex128{1 + 2i, 3 + 4i, 5 + 6i})  // -85 + 20i
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