// Package json provides high-performance JSON marshaling/unmarshaling using bytedance/sonic.
//
// This package wraps bytedance/sonic to provide a consistent JSON interface
// throughout the gophord library. Sonic provides significantly better performance
// than encoding/json through JIT compilation and SIMD acceleration.
//
// Usage:
//
//	data, err := json.Marshal(myStruct)
//	err := json.Unmarshal(data, &myStruct)
package json

import (
	stdjson "encoding/json"
	"reflect"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

// Marshal returns the JSON encoding of v using sonic.
// This is significantly faster than encoding/json.Marshal.
func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

// MarshalString returns the JSON encoding of v as a string using sonic.
func MarshalString(v interface{}) (string, error) {
	return sonic.MarshalString(v)
}

// MarshalIndent returns the indented JSON encoding of v.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	// sonic doesn't have MarshalIndent, use standard library
	return stdjson.MarshalIndent(v, prefix, indent)
}

// Unmarshal parses the JSON-encoded data and stores the result in v using sonic.
// This is significantly faster than encoding/json.Unmarshal.
func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

// UnmarshalString parses the JSON-encoded string and stores the result in v.
func UnmarshalString(s string, v interface{}) error {
	return sonic.UnmarshalString(s, v)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return sonic.Valid(data)
}

// Get extracts a value from JSON data using a path.
// This is useful for partial JSON parsing when you only need specific fields.
//
// Example:
//
//	node, err := json.Get(data, "user", "name")
//	name := node.String()
func Get(data []byte, path ...interface{}) (ast.Node, error) {
	return sonic.Get(data, path...)
}

// GetFromString extracts a value from a JSON string using a path.
func GetFromString(s string, path ...interface{}) (ast.Node, error) {
	return sonic.GetFromString(s, path...)
}

// ConfigDefault is the default sonic configuration.
var ConfigDefault = sonic.ConfigDefault

// ConfigStd provides encoding/json compatible behavior.
var ConfigStd = sonic.ConfigStd

// ConfigFastest provides the fastest encoding/decoding at the cost of some compatibility.
var ConfigFastest = sonic.ConfigFastest

// API provides a custom sonic API for advanced use cases.
type API = sonic.API

// Node represents a JSON value from partial parsing.
type Node = ast.Node

// Pretouch pre-compiles the encoder/decoder for a type.
// This is recommended for large types to improve first-call performance.
//
// Example:
//
//	func init() {
//		json.Pretouch(reflect.TypeOf(MyLargeStruct{}))
//	}
func Pretouch(t reflect.Type) error {
	return sonic.Pretouch(t)
}
