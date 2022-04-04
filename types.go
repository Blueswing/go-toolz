package toolz

// Signed signed integer
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned unsigned integer
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Numeric interface {
	Signed | Unsigned | Float | Complex
}

type OrderedNumeric interface {
	Signed | Unsigned | Float
}

type Ordered interface {
	Integer | Float | ~string
}

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}
