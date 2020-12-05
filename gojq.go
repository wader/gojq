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
// All functions with return type interface{} can return error on error
// array = []interface{}
// boolean = bool
// null = nil
// number = int | float64 | *big.Int
// object = map[string]interface{}
// string = string
// value = number | boolean | string | array | object | null
type JQValue interface {
	// JQValueLength is length of value, ex: value | length
	JQValueLength() interface{} // int

	// JQValueSliceLen slice length
	JQValueSliceLen() interface{} // int

	// JQValueIndex look up index for value, ex: value[index]
	// index -1 outside after slice, -2 outside before slice
	JQValueIndex(index int) interface{} // value

	// JQValueSlice slices value, ex: value[start:end]
	// start and end indexes are translated and clamped using JQValueSliceLen
	JQValueSlice(start int, end int) interface{} // []interface{}

	// JQValueKey look up key value of value: ex: value[name]
	JQValueKey(name string) interface{} // value

	// JQValueEach each of value, ex: value[]
	JQValueEach() interface{} // []PathValue

	// JQValueKeys keys for value, ex: value | keys
	JQValueKeys() interface{} // []string

	// JQValueHas check if value has key, ex: value | has(key)
	JQValueHas(key interface{}) interface{} // bool

	// JQValueType type of value, ex: value | type
	JQValueType() string // a JQType* constant

	// JQValueToNumber is value represented as a number, ex: value | tonumber
	JQValueToNumber() interface{} // number

	// JQValueToString is value represented as a string, ex: value | tostring
	JQValueToString() interface{} // string

	// JQValue value represented as a gojq compatible value
	JQValueToGoJQ() interface{} // value
}

// PathValue is a part of a jq path expression
type PathValue struct {
	Path, Value interface{}
}
