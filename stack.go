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

func (k *Stack) size() int {
	return len(*k)
}

func (k *Stack) empty() bool {
	return len(*k) == 0
}

func (k *Stack) push(s VType) {
	*k = append(*k, s)
}

func (k *Stack) pop() VType {
	if k.empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	*k = (*k)[:last]
	return s
}

func (k *Stack) popString() *VString {
	return NewVString(k.pop().String())
}

func (k *Stack) top() VType {
	if k.empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	return s
}

func (k *Stack) peek(n int) []VType {
	if k.size() <= n {
		return *k
	}
	last := len(*k)
	s := (*k)[last-n:last]
	return s
}

func (k *Stack) clear() {
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
	if k.size() < lim {
		str = "["
	}
	for i, elem := range k.peek(lim) {
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
