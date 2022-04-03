package defs

// UnaryFunc unary function
type UnaryFunc[InType any, OutType any] func(InType) OutType

// UnaryPred unary predicate
type UnaryPred[InType any] UnaryFunc[InType, bool]

// UnaryProc unary procedure
type UnaryProc[InType any] func(InType)

// BinaryFunc binary function
type BinaryFunc[InType1, InType2, OutType any] func(InType1, InType2) OutType

// BinaryPred binary predicate
type BinaryPred[InType1, InType2 any] BinaryFunc[InType1, InType2, bool]

// BinaryProc binary procedure
type BinaryProc[InType1, InType2 any] func(InType1, InType2)

type ReduceFunc[T any] BinaryFunc[T, T, T]

type PivotFunc[InType any, KeyType any, ValueType any] func(InType) (KeyType, ValueType)

type UnpivotFunc[KeyType any, ValueType any, OutType any] func(KeyType, ValueType) OutType
