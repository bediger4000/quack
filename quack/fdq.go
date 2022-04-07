package quack

import (
	"fmt"
	"io"
	"os"
)

type node struct {
	data interface{}
	next *node
}

type FdqQuack struct {
	B *node
	S *node
}

var _ Quack = (*FdqQuack)(nil)

func NewFdqQuack() Quack {
	fdq := &FdqQuack{}
	return fdq
}

func (fdq *FdqQuack) invariant() bool {
	lenS := 0
	for t := fdq.S; t != nil; t = t.next {
		lenS++
	}
	lenB := 0
	for t := fdq.B; t != nil; t = t.next {
		lenB++
	}

	return lenB >= lenS && lenS >= 1 && 3*lenS >= lenB
}

var z interface{}

func (fdq *FdqQuack) Push(x interface{}) {
	if !fdq.invariant() {
		fmt.Fprintf(os.Stderr, "push, invariant not met\n")
	}
}

func (fdq *FdqQuack) Pop() interface{} {
	if !fdq.invariant() {
		fmt.Fprintf(os.Stderr, "pop, invariant not met\n")
	}
	return z
}

func (fdq *FdqQuack) Pull() interface{} {
	if !fdq.invariant() {
		fmt.Fprintf(os.Stderr, "pull, invariant not met\n")
	}
	return z
}

func (fdq *FdqQuack) Print(out io.Writer) {
	fmt.Fprintf(out, "\n")
}
