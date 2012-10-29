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
// BUG: need to actually make this check for equality.
func (v *VList) Compare(other types.VType) int {
	return -2
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
