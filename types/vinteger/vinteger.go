// Package vinteger implements the integer type.
package vinteger

import (
	"strconv"

	"hawx.me/code/vodka/types"
)

// An integer, eg. `1`, `200`, etc.
type VInteger struct {
	value int
}

// String returns the value of the VInteger as a string.
func (v *VInteger) String() string {
	return strconv.FormatInt(int64(v.value), 10)
}

// Value returns the value of the VInteger.
func (v *VInteger) Value() interface{} {
	return v.value
}

// Type returns the name of the type, "integer".
func (v *VInteger) Type() string {
	return "integer"
}

func (v *VInteger) Next() types.Rangeable {
	return NewFromInt(v.value + 1)
}

func (v *VInteger) Prev() types.Rangeable {
	return NewFromInt(v.value - 1)
}

// Compare returns 0 if the values are equal, -1 if the value given is less, and
// 1 if the value given is greater.
func (v *VInteger) Compare(other types.VType) int {
	if val, same := other.(*VInteger); same {
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

func (v *VInteger) Copy() types.VType {
	return v
}

// New creates a new VInteger with the value parsed from the given string.
func New(s string) *VInteger {
	r := new(VInteger)
	r.value, _ = strconv.Atoi(s)
	return r
}

// NewFromInt creates a new VInteger with the value given.
func NewFromInt(i int) *VInteger {
	r := new(VInteger)
	r.value = i
	return r
}
