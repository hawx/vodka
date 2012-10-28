package vblock

import (
	"../../types"
	p "../../parser"
	"strings"
)

type VBlock struct {
	value string
}

func (v *VBlock) String() string {
	return "[ " + v.value + " ]"
}

func (v *VBlock) Value() interface{} {
	return p.Parse(v.value)
}

func (v *VBlock) BareValue() string {
	return v.value
}

func (v *VBlock) Type() string {
	return "block"
}

func (v *VBlock) Compare(other types.VType) int {
	if val, same := other.(*VBlock); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func New(s string) *VBlock {
	r := new(VBlock)
	r.value = strings.TrimSpace(s)
	return r
}
