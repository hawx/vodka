package main

import (
	"fmt"
	"strconv"
)

/*
 Synopsis:

 stk := NewStack()
 stk.push(1)
 stk.push(2)
 stk.push(3)
 fmt.Println(stk.String())
 //=> [1 2 3]

 stk.pop()
 fmt.Println(stk.String())
 //=> [1 2]

 stk.clear()
 fmt.Println(stk.String())
 //=> []
 */
type Stack []interface{}

func (k *Stack) size() int {
	return len(*k)
}

func (k *Stack) empty() bool {
	return len(*k) == 0
}

func (k *Stack) push(s interface{}) {
	*k = append(*k, s)
}

func (k *Stack) pop() interface{} {
	if k.empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
	*k = (*k)[:last]
	return s
}

func (k *Stack) popString() string {
	p := k.pop()

	if s, ok := p.(string); ok {
		return s
	} else if i, ok := p.(int); ok {
		return strconv.Itob(i, 10)
	}
	return "nil"
}

func (k *Stack) top() interface{} {
	if k.empty() {
		return nil
	}
	last := len(*k) - 1
	s := (*k)[last]
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

func NewStack() *Stack {
	return new(Stack)
}
