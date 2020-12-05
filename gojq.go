// Package gojq provides the parser and interpreter of gojq.
//
// Please refer to https://github.com/itchyny/gojq#usage-as-a-library for
// introduction of the usage as a library.
package gojq

// JQValue represents something that can be a jq value
type JQValue interface {
	// JQValueLength is length of value, ex: value | length
	JQValueLength() interface{}
	// JQValueIndex look up index for value, ex: value[index]
	JQValueIndex(index int) interface{}
	// JQValueSlice slices value, ex: value[start:end]
	JQValueSlice(start int, end int) interface{}
	// JQValueKey look up key value of value: ex: value[name]
	JQValueKey(name string) interface{}
	// JQValueEach each of value, ex: value[]
	// Should return []PathValue or error
	JQValueEach() interface{}
	// JQValueKeys keys for value, ex: value | keys
	JQValueKeys() interface{}
	// JQValueHasKey check if value has key, ex: value | has(key)
	JQValueHasKey(key interface{}) interface{}
	// JQValueType type of value, ex: value | type
	// Usually "array", "boolean", "object", "number", "string" or "null"
	JQValueType() string
	// JQValueToNumber is value represented as a number, ex: value | tonumber
	JQValueToNumber() interface{}
	// JQValueToString is value represented as a string, ex: value | tostring
	JQValueToString() interface{}
	// JQValue value represented as a gojq compatible value
	// Should be a int, float64, *big.Int, bool, string, []interface{}, map[string]interface{} or error
	JQValue() interface{}
}

// PathValue is a part of a jq path expression
type PathValue struct {
	Path, Value interface{}
}
