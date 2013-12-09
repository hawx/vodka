// Package vlist implements the list type.
package vlist

import (
	"github.com/hawx/vodka/types"
	"github.com/hawx/vodka/stack"
)

// A list type, eg. `(1 2 3)`.
type VList struct {
	value []types.VType
}

// String returns the value of the list formatted as a string.
func (v *VList) String() string {
	s := "("
	for i, item := range v.value {
		if i > 0 {
			s += " "
		}
		s += item.String()
	}
	return s + ")"
}

// Value returns the value of the list.
func (v *VList) Value() interface{} {
	return v.value
}

// Type returns the name of the type, "list".
func (v *VList) Type() string {
	return "list"
}

// Compare returns 0 if the values are equal, -2 otherwise.
func (v *VList) Compare(other types.VType) int {
	val, same := other.(*VList)
	if !same { return -2 }

	if len(v.value) != len(val.value) {
		return -2
	}

	for i := 0; i < len(v.value); i++ {
		if c := v.value[i].Compare(val.value[i]); c != 0 {
			return -2
		}
	}

	return 0
}

// New creates a new list from the Stack given.
func New(stk *stack.Stack) *VList {
	r := new(VList)
	r.value = *stk
	return r
}

// NewFromList creates a new list from the list given.
func NewFromList(list []types.VType) *VList {
	r := new(VList)
	r.value = list
	return r
}
