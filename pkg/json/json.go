package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Marshal returns the JSON encoding of v using encoding/json.
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalString returns the JSON encoding of v as a string using encoding/json.
func MarshalString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

// MarshalIndent returns the indented JSON encoding of v.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal parses the JSON-encoded data and stores the result in v using encoding/json.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// UnmarshalString parses the JSON-encoded string and stores the result in v.
func UnmarshalString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	var v interface{}
	return json.Unmarshal(data, &v) == nil
}

// Node is a placeholder for sonic.ast.Node compatibility.
type Node struct {
	val interface{}
}

func (n Node) String() string {
	if s, ok := n.val.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", n.val)
}

func (n Node) Int64() (int64, error) {
	switch v := n.val.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("not an integer: %T", n.val)
	}
}

func (n Node) Bool() (bool, error) {
	if b, ok := n.val.(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("not a boolean: %T", n.val)
}

func (n Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.val)
}

// Get extracts a value from JSON data using a path.
// This is a simplified shim for sonic.Get.
func Get(data []byte, path ...interface{}) (Node, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return Node{}, err
	}

	curr := v
	for _, p := range path {
		switch key := p.(type) {
		case string:
			if m, ok := curr.(map[string]interface{}); ok {
				curr = m[key]
			} else {
				return Node{}, fmt.Errorf("not a map at key %s", key)
			}
		case int:
			if a, ok := curr.([]interface{}); ok {
				if key >= 0 && key < len(a) {
					curr = a[key]
				} else {
					return Node{}, fmt.Errorf("index out of range: %d", key)
				}
			} else {
				return Node{}, fmt.Errorf("not a slice at index %d", key)
			}
		default:
			return Node{}, fmt.Errorf("unsupported path type: %T", p)
		}
	}

	return Node{val: curr}, nil
}

// GetFromString extracts a value from a JSON string using a path.
func GetFromString(s string, path ...interface{}) (Node, error) {
	return Get([]byte(s), path...)
}

// Pretouch is a no-op shim for sonic.Pretouch.
func Pretouch(t reflect.Type) error {
	return nil
}
