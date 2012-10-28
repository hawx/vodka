package vboolean

import "../../types"

type VBoolean struct {
	value bool
}

func (v *VBoolean) String() string {
	if v.value {
		return "true"
	}
	return "false"
}

func (v *VBoolean) Value() interface{} {
	return v.value
}

func (v *VBoolean) Type() string {
	return "boolean"
}

func (v VBoolean) Compare(other types.VType) int {
	if val, same := other.(*VBoolean); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func New(val bool) *VBoolean {
	b := new(VBoolean)
	b.value = val
	return b
}

func True() *VBoolean {
	return New(true)
}

func False() *VBoolean {
	return New(false)
}
