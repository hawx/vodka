// Package vrange implements the range type.
package vrange

import (
	"github.com/hawx/vodka/types"
	"github.com/hawx/vodka/types/vinteger"
	"github.com/hawx/vodka/types/vlist"
	"strings"
)

type VRange struct {
	start types.Rangeable
	end   types.Rangeable
}

func (v *VRange) String() string {
	return v.start.String() + ".." + v.end.String()
}

func (v *VRange) Value() interface{} {
	return []int{v.start.Value().(int), v.end.Value().(int)}
}

func (v *VRange) Type() string {
	return "range"
}

func (v *VRange) Compare(other types.VType) int {
	if o, ok := other.(*VRange); ok {
		if v.start.Compare(o.start) == 0 && v.end.Compare(o.end) == 0 {
			return 0
		}

		return -2
	}

	return -2
}

func (v *VRange) Max() types.Rangeable {
	if v.start.Compare(v.end) == 1 {
		return v.start
	}

	return v.end
}

func (v *VRange) Min() types.Rangeable {
	if v.start.Compare(v.end) == -1 {
		return v.start
	}

	return v.end
}

func (v *VRange) List() *vlist.VList {
	l := []types.VType{}

	if v.start.Compare(v.end) == -1 {
		for i := v.start; i.Compare(v.end) != 1; i = i.Next() {
			l = append(l, i)
		}
	} else {
		for i := v.start; i.Compare(v.end) != -1; i = i.Prev() {
			l = append(l, i)
		}
	}

	return vlist.NewFromList(l)
}

func New(s string) *VRange {
	parts := strings.Split(s, "..")

	r := new(VRange)
	r.start = vinteger.New(parts[0])
	r.end = vinteger.New(parts[1])
	return r
}

func NewFromStartAndEnd(start, end types.Rangeable) *VRange {
	return &VRange{
		start: start,
		end:   end,
	}
}
