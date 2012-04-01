package main

import (
	"strconv"
	"strings"
)

// TYPES --------------------------------------------

type VType interface {
	String()            string
	Value()             interface{}
	CompareWith(VType)  int
}

// SPECIALS ---------------------------------------------

type VNilType struct { }

func (v *VNilType) String() string {
	return "nil"
}

func (v *VNilType) Value() interface{} {
	return nil
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

// TOKENS ---------------------------------------------

type VTokens struct {
	value string
}

func (v *VTokens) String() string {
	return "tokens:[ " + v.value + " ]"
}

func (v *VTokens) Value() interface{} {
	return Parse(v.value)
}

func (v *VTokens) CompareWith(other VType) int {
	if val, same := other.(*VTokens); same {
		if val.value == v.value {
			return 0
		}
	}
	return -2
}

func NewVTokens(s string) *VTokens {
	r := new(VTokens)
	r.value = strings.TrimSpace(s)
	return r
}

// STRING ---------------------------------------------

type VString struct {
	value string
}

func (v *VString) String() string {
	return v.value
}

func (v *VString) Value() interface{} {
	return v.value
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
	return strconv.Itob(v.value, 10)
}

func (v *VInteger) Value() interface{} {
	return v.value
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
