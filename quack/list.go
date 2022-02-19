package quack

import "io"

type ListItem struct {
	data interface{}
	next *ListItem
	prev *ListItem
}

type ListQuack struct {
	head *ListItem
	tail *ListItem
}

var _ Quack = (*ListQuack)(nil)

func NewListQuack() Quack {
	lq := &ListQuack{
		head: &ListItem{},
		tail: &ListItem{},
	}
	lq.head.next = lq.tail
	lq.tail.prev = lq.head
	return lq
}

func (lq *ListQuack) Push(x interface{}) {
	n := &ListItem{data: x}

	n.next = lq.head.next
	n.prev = n.next.prev

	lq.head.next = n
	n.next.prev = n
}

func (lq *ListQuack) Pop() interface{} {
	if lq.head.next == lq.tail {
		return nil
	}
	n := lq.head.next
	return n.data
}

func (lq *ListQuack) Pull() interface{} {
	return nil
}

func (lq *ListQuack) Print(out io.Writer) {
}
