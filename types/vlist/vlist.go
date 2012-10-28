package vlist

import (
	"../../types"
	"../../stack"
)

type VList struct {
	value []types.VType
}

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

func (v *VList) Value() interface{} {
	return v.value
}

func (v *VList) Type() string {
	return "list"
}

func (v *VList) Compare(other types.VType) int {
	return -2
}

func New(stk *stack.Stack) *VList {
	r := new(VList)
	r.value = *stk
	return r
}

func NewFromList(list []types.VType) *VList {
	r := new(VList)
	r.value = list
	return r
}
