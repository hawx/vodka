package vnil

import "../../types"

type VNilType struct{}

func (v *VNilType) String() string {
	return "nil"
}

func (v *VNilType) Value() interface{} {
	return nil
}

func (v *VNilType) Type() string {
	return "nil"
}

func (v *VNilType) Compare(other types.VType) int {
	if _, same := (other).(*VNilType); same {
		return 0
	}
	return -2
}

func New() *VNilType {
	r := new(VNilType)
	return r
}
