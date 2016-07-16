// Package vdict implements the dictionary type.
package vdict

import (
	"hawx.me/code/vodka/stack"
	"hawx.me/code/vodka/types"
	"hawx.me/code/vodka/types/vboolean"
	"hawx.me/code/vodka/types/vnil"
)

// A dictionary type, eg. `{'a' -> 1, 'b' -> 2}`.
type VDict struct {
	keys   []types.VType
	values []types.VType
}

// String returns the value of the dictionary formatted as a string.
func (v *VDict) String() string {
	s := "{"
	for i := range v.keys {
		s += v.keys[i].String() + " " + v.values[i].String() + " "
	}
	return s[:len(s)-1] + "}"
}

// Value returns the value of the dictionary.
func (v *VDict) Value() interface{} {
	m := map[types.VType]types.VType{}
	for i := range v.keys {
		m[v.keys[i]] = v.values[i]
	}
	return m
}

func (v *VDict) Get(key types.VType) types.VType {
	for i, k := range v.keys {
		if k.Compare(key) == 0 {
			return v.values[i]
		}
	}

	return vnil.New()
}

func (v *VDict) Has(key types.VType) types.VType {
	for _, k := range v.keys {
		if k.Compare(key) == 0 {
			return vboolean.True()
		}
	}

	return vboolean.False()
}

func (v *VDict) Merge(other *VDict) types.VType {
	for i := range other.keys {
		if v.Has(other.keys[i]).Compare(vboolean.True()) == 0 {
			continue
		}

		v.keys = append(v.keys, other.keys[i])
		v.values = append(v.values, other.values[i])
	}

	return v
}

// Type returns the name of the type, "dict".
func (v *VDict) Type() string {
	return "dict"
}

// Compare returns 0 if the values are equal, -2 otherwise.
func (v *VDict) Compare(other types.VType) int {
	val, same := other.(*VDict)
	if !same {
		return -2
	}

	// check same size
	if len(v.keys) != len(val.keys) {
		return -2
	}

	// check has same keys, and values
	for _, k := range v.keys {
		if val.Has(k).Compare(vboolean.False()) == 0 {
			return -2
		}

		if v.Get(k).Compare(val.Get(k)) != 0 {
			return -2
		}
	}

	return 0
}

func (v *VDict) Copy() types.VType {
	return v
}

func New(stk *stack.Stack) *VDict {
	r := new(VDict)
	r.keys = []types.VType{}
	r.values = []types.VType{}

	for !stk.Empty() {
		v := stk.Pop()
		r.values = append(r.values, v)

		k := stk.Pop()
		r.keys = append(r.keys, k)
	}

	return r
}

func NewFromMap(m map[types.VType]types.VType) *VDict {
	r := new(VDict)
	r.keys = []types.VType{}
	r.values = []types.VType{}

	for k, v := range m {
		r.values = append(r.values, v)
		r.keys = append(r.keys, k)
	}

	return r
}
