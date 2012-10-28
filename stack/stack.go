// Package stack implements a stack of VType values.
package stack

import (
	"fmt"
	"../types"
	"../types/vstring"
)


type Stack []types.VType

// Size returns the size of the Stack.
func (k *Stack) Size() int {
	return len(*k)
}

// Empty returns true if the Stack contains no items, false otherwise.
func (k *Stack) Empty() bool {
	return k.Size() == 0
}

// Push adds a new item to the top of the Stack.
func (k *Stack) Push(s types.VType) {
	*k = append(*k, s)
}

// Pop takes the top element from the Stack and returns it.
func (k *Stack) Pop() types.VType {
	if k.Empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	*k = (*k)[:last]
	return s
}

// PopString takes the top element from the Stack and returns the string value
// of it.
func (k *Stack) PopString() *vstring.VString {
	return vstring.New(k.Pop().String())
}

// Top returns the top element of the Stack without removing it.
func (k *Stack) Top() types.VType {
	if k.Empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	return s
}

// Peek returns a list of the top n items on the Stack.
func (k *Stack) Peek(n int) []types.VType {
	if k.Size() <= n {
		return *k
	}
	last := len(*k)
	s := (*k)[last-n : last]
	return s
}

// Clear removes all elements from the Stack.
func (k *Stack) Clear() {
	*k = (*k)[0:0]
}

// String returns a string representation of the Stack.
func (k Stack) String() string {
	str := "["
	for i, elem := range k {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// TruncatedString returns a string representation of the Stack showing only the
// top 7 (or less) values.
func (k Stack) TruncatedString() string {
	lim := 7
	str := "... "
	if k.Size() < lim {
		str = "["
	}
	for i, elem := range k.Peek(lim) {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// New returns a new empty Stack.
func New() *Stack {
	return new(Stack)
}
