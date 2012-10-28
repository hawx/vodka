package vstring

import "../../types"

type VString struct {
	value string
}

func (v *VString) String() string {
	return "'" + v.value + "'"
}

func (v *VString) Value() interface{} {
	return v.value
}

func (v *VString) Type() string {
	return "string"
}

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

func New(s string) *VString {
	r := new(VString)
	r.value = s
	return r
}
