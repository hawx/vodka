// Package vnil implements the nil type.
package vnil

import "hawx.me/code/vodka/types"

// the nil value, `nil`.
type VNil struct{}

// String returns the string value of the VNil, "nil".
func (v *VNil) String() string {
	return "nil"
}

// Value returns the value of the VNil, `nil`.
func (v *VNil) Value() interface{} {
	return nil
}

// Type returns the name of the type, "nil".
func (v *VNil) Type() string {
	return "nil"
}

// Compare returns 0 if the other value is a VNil, and -2 otherwise.
func (v *VNil) Compare(other types.VType) int {
	if _, same := (other).(*VNil); same {
		return 0
	}
	return -2
}

func (v *VNil) Copy() types.VType {
	return v
}

// New creates a new VNil
func New() *VNil {
	return &VNil{}
}
