package quack

import (
	"fmt"
	"io"
)

type StackItem struct {
	data interface{}
	next *StackItem
}

func (si *StackItem) String() string {
	return fmt.Sprintf("%v", si.data)
}

type StackQuack struct {
	pushstack *StackItem
	pullstack *StackItem
}

var _ Quack = (*StackQuack)(nil)

func NewStackQuack() Quack {
	return &StackQuack{}
}

func (sq *StackQuack) Push(x interface{}) {
	n := &StackItem{data: x}
	n.next = sq.pushstack
	sq.pushstack = n
}

func (sq *StackQuack) Pop() interface{} {
	top := sq.pushstack
	if top != nil {
		sq.pushstack = top.next
		return top.data
	}

	if sq.pullstack != nil {
		for n := sq.pullstack; n != nil; n = sq.pullstack {
			sq.pullstack = n.next
			n.next = sq.pushstack
			sq.pushstack = n
		}
		top := sq.pushstack
		if top != nil {
			sq.pushstack = top.next
			return top.data
		}
	}

	return nil
}

func (sq *StackQuack) Pull() interface{} {
	pull := sq.pullstack
	if pull != nil {
		sq.pullstack = pull.next
		return pull.data
	}
	// pullstack is empty
	if sq.pushstack == nil {
		return nil
	}
	for n := sq.pushstack; n != nil; n = sq.pushstack {
		sq.pushstack = n.next
		n.next = sq.pullstack
		sq.pullstack = n
	}
	bottom := sq.pullstack
	if bottom != nil {
		sq.pullstack = bottom.next
		return bottom.data
	}
	return nil
}

func (sq *StackQuack) Print(out io.Writer) {
	fmt.Fprintf(out, "Push stack: ")
	for n := sq.pushstack; n != nil; n = n.next {
		fmt.Fprintf(out, "%v -> ", n)
	}
	fmt.Fprintf(out, "\nPull stack: ")
	for n := sq.pullstack; n != nil; n = n.next {
		fmt.Fprintf(out, "%v -> ", n)
	}
	fmt.Fprintf(out, "\n")
}
