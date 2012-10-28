package vinteger

import (
	"../../types"
	"strconv"
)

type VInteger struct {
	value int
}

func (v *VInteger) String() string {
	return strconv.FormatInt(int64(v.value), 10)
}

func (v *VInteger) Value() interface{} {
	return v.value
}

func (v *VInteger) Type() string {
	return "integer"
}

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

func New(s string) *VInteger {
	r := new(VInteger)
	r.value, _ = strconv.Atoi(s)
	return r
}

func NewFromInt(i int) *VInteger {
	r := new(VInteger)
	r.value = i
	return r
}
