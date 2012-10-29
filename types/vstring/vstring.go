// Package vstring implements the string type.
package vstring

import "../../types"

// A string value, eg. `"hello world"`.
type VString struct {
	value string
}

// String returns the value formatted as a string.
func (v *VString) String() string {
	return "'" + v.value + "'"
}

// Value returns the value.
func (v *VString) Value() interface{} {
	return v.value
}

// Type returns the name of the type, "string".
func (v *VString) Type() string {
	return "string"
}

// Compare returns 0 if the other string is equal, -1 if it less than or 1 if it
// greater than.
func (v *VString) Compare(other types.VType) int {
	if val, same := other.(*VString); same {
		if val.value == v.value {
			return 0
		} else if val.value < v.value {
			return 1
		} else {
			return -1
		}
	}
	return -2
}

// New creates a new VString with the value given.
func New(s string) *VString {
	r := new(VString)
	r.value = s
	return r
}
