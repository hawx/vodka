package main

import (
	"strconv"
	"strings"
)

// TYPES --------------------------------------------

type VType interface {
	String() string
	Value() interface{}
	Type() string
	CompareWith(VType) int
}

// SPECIALS ---------------------------------------------

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

func (v *VNilType) CompareWith(other VType) int {
	if _, same := (other).(*VNilType); same {
		return 0
	}
	return -2
}

func VNil() *VNilType {
	r := new(VNilType)
	return r
}

// BOOLEAN ---------------------------------------------

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

func (v VBoolean) CompareWith(other VType) int {
	if val, same := other.(*VBoolean); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func NewVBoolean(val bool) *VBoolean {
	b := new(VBoolean)
	b.value = val
	return b
}

func VTrue() *VBoolean {
	return NewVBoolean(true)
}

func VFalse() *VBoolean {
	return NewVBoolean(false)
}

// BLOCK ---------------------------------------------

type VBlock struct {
	value string
}

func (v *VBlock) String() string {
	return "block:[ " + v.value + " ]"
}

func (v *VBlock) Value() interface{} {
	return Parse(v.value)
}

func (v *VBlock) Type() string {
	return "block"
}

func (v *VBlock) CompareWith(other VType) int {
	if val, same := other.(*VBlock); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func NewVBlock(s string) *VBlock {
	r := new(VBlock)
	r.value = strings.TrimSpace(s)
	return r
}

// STRING ---------------------------------------------

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

func (v *VString) CompareWith(other VType) int {
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

func NewVString(s string) *VString {
	r := new(VString)
	r.value = s
	return r
}

// Integer ------------------------------------------------

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

func (v *VInteger) CompareWith(other VType) int {
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

func NewVInteger(s string) *VInteger {
	r := new(VInteger)
	r.value, _ = strconv.Atoi(s)
	return r
}

func NewVIntegerInt(i int) *VInteger {
	r := new(VInteger)
	r.value = i
	return r
}

// List ------------------------------------------------------

type VList struct {
	value []VType
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

func (v *VList) CompareWith(other VType) int {
	return -2
}

func NewVList(stk *Stack) *VList {
	r := new(VList)
	r.value = *stk
	return r
}

func NewVListList(list []VType) *VList {
	r := new(VList)
	r.value = list
	return r
}
