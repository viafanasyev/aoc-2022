package utils

import (
    "fmt"
)

type Stack[T any] struct {
    top *Node[T]
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T] { nil }
}

func (st *Stack[T]) Push(value T) {
    curNode := &Node[T] { value, st.top }
    st.top = curNode
}

func (st *Stack[T]) Peek() T {
    if st.top == nil {
        panic(fmt.Sprintf("Trying to peek in empty stack"))
    }
    return st.top.value
}

func (st *Stack[T]) Pop() T {
    if st.top == nil {
        panic(fmt.Sprintf("Trying to pop from empty stack"))
    }
    value := st.top.value
    st.top = st.top.next
    return value
}
