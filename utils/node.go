package utils

type Node[T any] struct {
    value T
    next *Node[T]
}
