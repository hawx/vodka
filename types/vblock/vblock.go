// Package vblock implements the block type.
package vblock

import (
	"strings"

	"hawx.me/code/vodka/parser"
	"hawx.me/code/vodka/types"
)

// A block of vodka code, eg. `[mult]` or `:add`.
type VBlock struct {
	value string
}

// String returns the string value of the VBlock.
func (v *VBlock) String() string {
	return "[ " + v.value + " ]"
}

// Value returns the parsed value of the VBlock.
func (v *VBlock) Value() interface{} {
	return parser.Parse(v.value)
}

// BareValue returns the string contained by the VBlock.
func (v *VBlock) BareValue() string {
	return v.value
}

// Type returns the type of a VBlock, which is "block".
func (v *VBlock) Type() string {
	return "block"
}

// Compare returns 0 if other matches the VBlock being called on, or -2 if not
// equal.
func (v *VBlock) Compare(other types.VType) int {
	if val, same := other.(*VBlock); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func (v *VBlock) Copy() types.VType {
	return v
}

// New creates a new VBlock with the string given.
func New(s string) *VBlock {
	return &VBlock{
		value: strings.TrimSpace(s),
	}
}
