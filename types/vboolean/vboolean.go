// Package vboolean implements the boolean type.
package vboolean

import "../../types"

// A boolean value, either `true` or `false`.
type VBoolean struct {
	value bool
}

// String returns the string value of the VBoolean.
func (v *VBoolean) String() string {
	if v.value {
		return "true"
	}
	return "false"
}

// Value returns the boolean value.
func (v *VBoolean) Value() interface{} {
	return v.value
}

// Type returns the name of the type, "boolean".
func (v *VBoolean) Type() string {
	return "boolean"
}

// Compare returns 0 if the values are equal, -2 otherwise.
func (v VBoolean) Compare(other types.VType) int {
	if val, same := other.(*VBoolean); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

// New creates a new VBoolean with the value given.
func New(val bool) *VBoolean {
	b := new(VBoolean)
	b.value = val
	return b
}

// True returns a new true VBoolean.
func True() *VBoolean {
	return New(true)
}

// False returns a new false VBoolean.
func False() *VBoolean {
	return New(false)
}
