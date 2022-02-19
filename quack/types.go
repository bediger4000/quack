package quack

import "io"

type Quack interface {
	Push(interface{})
	Pop() interface{}
	Pull() interface{}
	Print(io.Writer)
}
