// Package gojq provides the parser and the interpreter of gojq.
// Please refer to [Usage as a library] for introduction.
//
// [Usage as a library]: https://github.com/itchyny/gojq#usage-as-a-library
package gojq

// TODO: use in TypeOf?
const (
	JQTypeArray   = "array"
	JQTypeBoolean = "boolean"
	JQTypeNull    = "null"
	JQTypeNumber  = "number"
	JQTypeObject  = "object"
	JQTypeString  = "string"
)

// JQValue represents something that can be a jq value
// All functions with return type any can return error on error
// array = []any
// boolean = bool
// null = nil
// number = int | float64 | *big.Int
// object = map[string]any
// string = string
// value = number | boolean | string | array | object | null
type JQValue interface {
	// JQValueLength is length of value, ex: value | length
	JQValueLength() any // int

	// JQValueSliceLen slice length
	JQValueSliceLen() any // int

	// JQValueIndex look up index for value, ex: value[index]
	// index -1 outside after slice, -2 outside before slice
	JQValueIndex(index int) any // value

	// JQValueSlice slices value, ex: value[start:end]
	// start and end indexes are translated and clamped using JQValueSliceLen
	JQValueSlice(start int, end int) any // []any

	// JQValueKey look up key value of value: ex: value[name]
	JQValueKey(name string) any // value

	// JQValueEach each of value, ex: value[]
	JQValueEach() any // []PathValue

	// JQValueKeys keys for value, ex: value | keys
	JQValueKeys() any // []string

	// JQValueHas check if value has key, ex: value | has(key)
	JQValueHas(key any) any // bool

	// JQValueType type of value, ex: value | type
	JQValueType() string // a JQType* constant

	// JQValueToNumber is value represented as a number, ex: value | tonumber
	JQValueToNumber() any // number

	// JQValueToString is value represented as a string, ex: value | tostring
	JQValueToString() any // string

	// JQValue value represented as a gojq compatible value
	JQValueToGoJQ() any // value
}

// PathValue is a part of a jq path expression
type PathValue struct {
	Path, Value any
}
