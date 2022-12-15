package utils

import (
    "fmt"
)

type Queue[T any] struct {
    front *Node[T]
    back *Node[T]
}

func NewQueue[T any]() *Queue[T] {
    return &Queue[T] { nil, nil }
}

func (q *Queue[T]) Enqueue(value T) {
    oldBack := q.back
    q.back = &Node[T] { value, nil }
    if q.front == nil {
        q.front = q.back
    } else {
        oldBack.next = q.back
    }
}

func (q *Queue[T]) Peek() T {
    if q.front == nil {
        panic(fmt.Sprintf("Trying to peek in empty queue"))
    }
    return q.front.value
}

func (q *Queue[T]) Dequeue() T {
    if q.front == nil {
        panic(fmt.Sprintf("Trying to pop from empty queue"))
    }
    value := q.front.value
    q.front = q.front.next
    if q.front == nil {
        q.back = nil
    }
    return value
}

func (q *Queue[T]) IsEmpty() bool {
    return q.front == nil
}
