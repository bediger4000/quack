package quack

type StackItem struct {
	data interface{}
	next *StackItem
}

type StackQuack struct {
	pushstack  *StackItem
	pullstack  *StackItem
	sparestack *StackItem
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

	if sq.pushstack != nil {
		for n := sq.pushstack; n != nil; n = sq.pushstack {
			n.next = sq.pullstack
			sq.pullstack = n
		}
		bottom := sq.pullstack
		if bottom != nil {
			sq.pullstack = bottom.next
			return bottom.data
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
