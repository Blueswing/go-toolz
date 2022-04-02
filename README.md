# go-toolz
A functional library written in Go.

## Examples

### GroupBy
```go
seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
res := GroupBy(seq, func(x int) int { return x % 2 })
fmt.Println(res)
```

Output:
```
map[0:[2 4 6 8] 1:[1 3 5 7]]
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
	res := Pivot(s, func(x map[string]interface{}) (int, string) { return x["k"].(int), x["v"].(string) })
	fmt.Println(res)
```

Output:
```
map[1:one 2:two 3:three 4:four 5:five]
```
