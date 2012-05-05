package main

import (
	"fmt"
)

// Synopsis:
//
//   stk := NewStack()
//   stk.push(1)
//   stk.push(2)
//   stk.push(3)
//   fmt.Println(stk.String())
//   ;=> [1 2 3]
//
//   stk.pop()
//   fmt.Println(stk.String())
//   ;=> [1 2]
//
//   stk.clear()
//   fmt.Println(stk.String())
//   ;=> []

type Stack []VType

func (k *Stack) Size() int {
	return len(*k)
}

func (k *Stack) Empty() bool {
	return len(*k) == 0
}

func (k *Stack) Push(s VType) {
	*k = append(*k, s)
}

func (k *Stack) Pop() VType {
	if k.Empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	*k = (*k)[:last]
	return s
}

func (k *Stack) PopString() *VString {
	return NewVString(k.Pop().String())
}

func (k *Stack) Top() VType {
	if k.Empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	return s
}

func (k *Stack) Peek(n int) []VType {
	if k.Size() <= n {
		return *k
	}
	last := len(*k)
	s := (*k)[last-n : last]
	return s
}

func (k *Stack) Clear() {
	*k = (*k)[0:0]
}

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

func NewStack() *Stack {
	return new(Stack)
}
