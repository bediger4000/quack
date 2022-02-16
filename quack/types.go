package quack

type Quack interface {
	Push(interface{})
	Pop() interface{}
	Pull() interface{}
}
